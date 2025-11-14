package repository

import (
	"context"
	"shorten-url/package/database/mysql"
)

type repository struct {
	db *mysql.Client
}

func (r *repository) GetMaxShortenCode(context.Context) (int, error) {
	panic("unimplemented")
}

func NewRepository(db *mysql.Client) *repository {
	return &repository{
		db: db,
	}
}
