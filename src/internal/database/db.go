package database

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/migrate"
	"github.com/hawa130/computility-cloud/ent/privacy"
	_ "github.com/lib/pq"
)

var client *ent.Client
var AllowContext = privacy.DecisionContext(context.Background(), privacy.Allow)

func Open(driverName, dataSourceName string) (*ent.Client, error) {
	var err error
	client, err = ent.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
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
