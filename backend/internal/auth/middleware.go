package auth

import (
	"context"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/hawa130/serverx/config"
	"github.com/hawa130/serverx/ent"
	"github.com/hawa130/serverx/ent/user"
	"github.com/hawa130/serverx/internal/database"
	"github.com/labstack/echo/v4"
)

func getUserFromToken(token string) (*ent.User, *JWTClaims, error) {
	claims, err := ParseToken(token)
	if err != nil {
		return nil, nil, err
	}

	requestUser, err := database.Client().User.Query().
		Where(user.IDEQ(claims.Subject)).
		Only(database.AllowContext)
	if err != nil {
		return nil, nil, err
	}

	return requestUser, claims, nil
}

// Middleware is a middleware for renewing JWT tokens and setting user in context
func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return next(c)
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == "" {
				return next(c)
			}

			requestUser, claims, err := getUserFromToken(tokenString)
			if err != nil {
				return next(c)
			}

			if time.Until(time.Unix(claims.ExpiresAt, 0)) < config.Config().JWT.RenewDuration*time.Hour {
				newToken, err := GenerateToken(requestUser.ID)
				if err != nil {
					return next(c)
				}
				c.Response().Header().Set("X-Set-Token", newToken)
			}

			ctx := context.WithValue(c.Request().Context(), "user", requestUser)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

func FromContext(c context.Context) (*ent.User, bool) {
	u, ok := c.Value("user").(*ent.User)
	return u, ok
}

func WebsocketInit(ctx context.Context, initPayload *transport.InitPayload) (context.Context, *transport.InitPayload, error) {
	token := initPayload.GetString("token")
	if token == "" {
		return ctx, initPayload, nil
	}

	requestUser, _, err := getUserFromToken(token)
	if err != nil {
		return ctx, initPayload, nil
	}

	ctx = context.WithValue(ctx, "user", requestUser)
	return ctx, initPayload, nil
}
