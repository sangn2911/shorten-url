package port

import (
	"context"
)

type Service interface {
	Encode(context.Context, string) (string, error)
	Decode(context.Context, string) (string, error)
}
