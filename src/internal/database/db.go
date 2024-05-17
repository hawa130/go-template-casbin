package database

import (
	"context"

	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
)

var client *ent.Client
var isOpen = false

func Client(driverName, dataSourceName string) (*ent.Client, error) {
	if isOpen {
		return client, nil
	}

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

	isOpen = true
	return client, nil
}

func Close() error {
	if client != nil {
		if err := client.Close(); err != nil {
			return err
		}
		isOpen = false
	}
	return nil
}
