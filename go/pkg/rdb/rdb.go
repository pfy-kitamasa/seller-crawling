package rdb

import (
	"context"
	"fmt"

	"github.com/panforyou/seller-finding/go/ent"
	"github.com/panforyou/seller-finding/go/ent/migrate"
	"golang.org/x/xerrors"
)

func NewClient(host string, port int, name, user, password string) (*ent.Client, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		host, port, name, user, password)
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, xerrors.Errorf("failed to connect db: %v", err)
	}

	return client, nil
}

func Migration(ctx context.Context, client *ent.Client) error {
	if err := client.Debug().Schema.Create(ctx,
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	); err != nil {
		return xerrors.Errorf("failed to migrage db schema: %v", err)
	}

	return nil
}
func MigrationWithSeed(ctx context.Context, client *ent.Client, fn func() error) error {
	if err := client.Debug().Schema.Create(ctx,
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	); err != nil {
		return xerrors.Errorf("failed to migrage db schema: %v", err)
	}

	return fn()
}
