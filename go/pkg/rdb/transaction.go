package rdb

import (
	"context"

	"github.com/panforyou/seller-finding/go/ent"

	"golang.org/x/xerrors"
)

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return xerrors.Errorf("rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return xerrors.Errorf("committing transaction: %v", err)
	}
	return nil
}
