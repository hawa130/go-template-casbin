package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hawa130/serverx/config"
	_ "github.com/hawa130/serverx/ent/runtime"
	"github.com/hawa130/serverx/graph"
	"github.com/hawa130/serverx/internal/auth"
	"github.com/hawa130/serverx/internal/database"
	"github.com/hawa130/serverx/internal/logger"
	"github.com/hawa130/serverx/internal/perm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/xid"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	return &Server{}
}

func (r *Server) HandleGraphql() {
	client := database.Client()
	cfg := config.Config()

	r.echo.POST(cfg.GraphQL.Endpoint, func(c echo.Context) error {
		srv := handler.New(graph.NewSchema(client))
		srv.Use(entgql.Transactioner{TxOpener: client})

		srv.AddTransport(transport.POST{})
		srv.AddTransport(transport.Websocket{
			KeepAlivePingInterval: 10 * time.Second,
			InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, *transport.InitPayload, error) {
				return auth.WebsocketInit(ctx, &initPayload)
			},
		})

		if cfg.GraphQL.Introspection {
			srv.Use(extension.Introspection{})
		}

		srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
			c.Set("operation_context", graphql.GetOperationContext(ctx))
			return next(ctx)
		})

		srv.ServeHTTP(c.Response(), c.Request())

		return nil
	})

	if cfg.GraphQL.Playground {
		r.echo.GET(
			cfg.GraphQL.PlaygroundEndpoint,
			echo.WrapHandler(playground.Handler("GraphQL playground", cfg.GraphQL.Endpoint)),
		)
	}
}

func (r *Server) Start() {
	log.Println("starting server")

	cfg := config.Config()

	err := logger.Init()
	if err != nil {
		log.Fatal("logger initialization error: ", err)
	}

	err = perm.Init(cfg.Database.Driver, cfg.Database.Url)
	if err != nil {
		log.Fatal("casbin initialization error: ", err)
	}

	_, err = database.Open(cfg.Database.Driver, cfg.Database.Url)
	if err != nil {
		log.Fatal("database initialization error: ", err)
	}

	r.echo = echo.New()
	r.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     cfg.Server.CORS.AllowedOrigins,
		AllowMethods:     cfg.Server.CORS.AllowedMethods,
		AllowHeaders:     cfg.Server.CORS.AllowedHeaders,
		ExposeHeaders:    cfg.Server.CORS.ExposedHeaders,
		AllowCredentials: cfg.Server.CORS.AllowCredentials,
		MaxAge:           cfg.Server.CORS.MaxAge,
	}))
	r.echo.Use(middleware.Recover())
	r.echo.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return xid.New().String()
		},
	}))
	r.echo.Use(logger.Middleware())
	r.echo.Use(auth.Middleware())

	r.HandleGraphql()

	go func() {
		if err := r.echo.Start(cfg.Server.Address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server start failed: %v", err)
		}
	}()
}

func (r *Server) Stop() {
	log.Println("stopping server")

	if r.echo != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := r.echo.Shutdown(ctx); err != nil {
			log.Fatalf("server shutdown failed: %v", err)
		}
	}

	if err := database.Close(); err != nil {
		log.Fatalf("database close failed: %v", err)
	}

	logger.Sync()
	log.Println("server stopped")
}

func (r *Server) Restart() {
	log.Println("restarting server")
	r.Stop()
	r.Start()
}
