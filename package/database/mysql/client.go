package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"shorten-url/package/database"
	"shorten-url/package/logger"

	_ "github.com/go-sql-driver/mysql"
)

const mySQLTxKey = database.CtxTx("mysql_tx")

type Client struct {
	*sql.DB
}

func NewClient(cfg MySQLConfig) (*Client, error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	logger.Info("Connect to MySQL successfully")
	return &Client{
		DB: db,
	}, nil
}

func (c *Client) GetDBTX(ctx context.Context) database.DBTX {
	if tx := database.GetTx(ctx, mySQLTxKey); tx != nil {
		return tx
	}
	return c.DB
}

func (c *Client) Begin(ctx context.Context) (context.Context, error) {
	tx, err := c.DB.Begin()
	if err != nil {
		return nil, err
	}
	return database.SetTx(ctx, mySQLTxKey, tx), nil
}

func (c *Client) Commit(ctx context.Context) error {
	if tx := database.GetTx(ctx, mySQLTxKey); tx != nil {
		return tx.Commit()
	}
	return nil
}

func (c *Client) RollBack(ctx context.Context) error {
	if tx := database.GetTx(ctx, mySQLTxKey); tx != nil {
		return tx.Rollback()
	}
	return nil
}
