package logger

import (
	"context"
)

type CtxPID string

var ctxPID = CtxPID("pid")

func SetPID(ctx context.Context, pid string) context.Context {
	return context.WithValue(ctx, ctxPID, pid)
}

func GetPID(ctx context.Context) string {
	value := ctx.Value(ctxPID)
	if pid, ok := value.(string); ok {
		return pid
	}
	return ""
}
