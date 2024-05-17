package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hawa130/computility-cloud/config"
	_ "github.com/hawa130/computility-cloud/ent/runtime"
	"github.com/hawa130/computility-cloud/graph"
	"github.com/hawa130/computility-cloud/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

func startServer() {
	log.Println("starting server")

	cfg := config.GetConfig()

	c, err := database.Client(dialect.SQLite, cfg.Database.Url)
	if err != nil {
		log.Fatal("database initialization error: ", err)
	}

	e = echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	srv := handler.NewDefaultServer(graph.NewSchema(c))

	e.POST(cfg.GraphQL.EndPoint, echo.WrapHandler(srv))
	if cfg.GraphQL.Playground {
		e.GET(
			cfg.GraphQL.PlaygroundEndpoint,
			echo.WrapHandler(playground.Handler("GraphQL playground", cfg.GraphQL.EndPoint)),
		)
	}

	go func() {
		if err := e.Start(cfg.Server.Address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server start failed: %v", err)
		}
	}()
}

func stopServer() {
	log.Println("stopping server")

	if e != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			log.Fatalf("server shutdown failed: %v", err)
		}
	}

	if err := database.Close(); err != nil {
		log.Fatalf("database close failed: %v", err)
	}

	log.Println("server stopped")
}

func restartServer() {
	log.Println("restarting server")
	stopServer()
	startServer()
}
