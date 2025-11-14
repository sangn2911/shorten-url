package mysql

import (
	"context"
	"database/sql"
)

type CtxTx string

const ctxTx = CtxTx("mysql_tx")

func SetTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, ctxTx, tx)
}

func GetTx(ctx context.Context) *sql.Tx {
	value := ctx.Value(ctxTx)
	if tx, ok := value.(*sql.Tx); ok {
		return tx
	}
	return nil
}
