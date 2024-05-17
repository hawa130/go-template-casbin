package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hawa130/computility-cloud/config"
	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/migrate"
	_ "github.com/hawa130/computility-cloud/ent/runtime"
	"github.com/hawa130/computility-cloud/graph"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.GetConfig()

	client, err := ent.Open(dialect.SQLite, cfg.Database.Url)
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("running schema migration", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	srv := handler.NewDefaultServer(graph.NewSchema(client))

	e.POST("/query", echo.WrapHandler(srv))
	e.GET("/", echo.WrapHandler(playground.Handler("GraphQL playground", "/query")))

	if err := e.Start(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		e.Logger.Fatal(err)
	}
}
