package database

import (
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/hawa130/serverx/ent"
	"github.com/hawa130/serverx/ent/migrate"
	"github.com/hawa130/serverx/ent/privacy"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var client *ent.Client

var AllowContext = privacy.DecisionContext(context.Background(), privacy.Allow)

func Open(dataSourceName string) (*ent.Client, error) {
	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	client = ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropIndex(true),  // 迁移时删除不使用的索引，早期开发建议启用
		migrate.WithDropColumn(true), // 迁移时删除不使用的列，早期开发建议启用
	); err != nil {
		return nil, err
	}

	if err := seedData(AllowContext); err != nil {
		return nil, err
	}

	return client, nil
}

func Close() error {
	if client != nil {
		if err := client.Close(); err != nil {
			return err
		}
		client = nil
	}
	return nil
}

func Client() *ent.Client {
	return client
}
