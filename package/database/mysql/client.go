package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"shorten-url/package/database"
	"shorten-url/package/logger"

	_ "github.com/go-sql-driver/mysql"
)

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
	logger.Info("Connect to MySQL successfully")
	return &Client{
		DB: db,
	}, nil
}

func (c *Client) GetDBTX(ctx context.Context) database.DBTX {
	if tx := GetTx(ctx); tx != nil {
		return tx
	}
	return c.DB
}
