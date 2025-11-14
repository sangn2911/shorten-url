package port

import "context"

type Repository interface {
	GetMaxShortenCode(context.Context) (int, error)
}
