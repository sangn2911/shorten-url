package database

import (
	"context"
	"database/sql"
)

type CtxTx string

func SetTx(ctx context.Context, key CtxTx, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, key, tx)
}

func GetTx(ctx context.Context, key CtxTx) *sql.Tx {
	value := ctx.Value(key)
	if tx, ok := value.(*sql.Tx); ok {
		return tx
	}
	return nil
}
