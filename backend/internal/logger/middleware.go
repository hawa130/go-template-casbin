package logger

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func Middleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogMethod:    true,
		LogStatus:    true,
		LogUserAgent: true,
		LogLatency:   true,
		LogRemoteIP:  true,
		LogRequestID: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fields := []zap.Field{
				zap.String("start_time", v.StartTime.Format(time.RFC3339Nano)),
				zap.String("id", v.RequestID),
				zap.String("remote_ip", v.RemoteIP),
				zap.String("user_agent", v.UserAgent),
				zap.Duration("latency", v.Latency),
				zap.String("latency_human", v.Latency.String()),
			}

			restfulFields := append(fields,
				zap.String("method", v.Method),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
			)

			if v.URI == "/graphql" {
				oc, ok := c.Get("operation_context").(*graphql.OperationContext)
				if !ok {
					logger.Info("request", restfulFields...)
					return nil
				}
				graphqlFields := append(fields,
					zap.String("operation", oc.OperationName),
				)
				logger.Info("query", graphqlFields...)
				return nil
			}
			logger.Info("request", restfulFields...)
			return nil
		},
	})
}
