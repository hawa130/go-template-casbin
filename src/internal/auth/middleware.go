package auth

import (
	"context"
	"strings"
	"time"

	"github.com/hawa130/computility-cloud/config"
	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/internal/database"
	"github.com/labstack/echo/v4"
)

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

			claims, err := ParseToken(tokenString)
			if err != nil {
				return next(c)
			}

			user, err := database.Client().User.Get(context.Background(), claims.Uid)
			if err != nil {
				return next(c)
			}

			if time.Until(time.Unix(claims.ExpiresAt, 0)) < config.GetConfig().JWT.RenewDuration*time.Hour {
				newToken, err := GenerateToken(user.ID)
				if err != nil {
					return next(c)
				}
				c.Response().Header().Set("X-Set-Authorization", "Bearer "+newToken)
			}

			ctx := context.WithValue(c.Request().Context(), "user", user)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

func FromContext(c context.Context) (*ent.User, bool) {
	user, ok := c.Value("user").(*ent.User)
	return user, ok
}
