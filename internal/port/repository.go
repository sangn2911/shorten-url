package port

import (
	"context"
	"shorten-url/internal/model"
)

type Repository interface {
	Begin(context.Context) (context.Context, error)
	Commit(context.Context) error
	RollBack(context.Context) error
	GetUrlEncodeByOriginalUrl(context.Context, string) (model.UrlEncodeModel, error)
	GetLatestUrlEncode(context.Context) (model.UrlEncodeModel, error)
	InsertUrlEncode(context.Context, model.UrlEncodeModel) error
	GetUrlEncodeByShortenId(context.Context, string) (model.UrlEncodeModel, error)
}
