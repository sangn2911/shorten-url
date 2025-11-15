package middleware

import (
	"net/http"
	"shorten-url/package/logger"

	"github.com/google/uuid"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pid := uuid.NewString()
		ctx := logger.SetPID(r.Context(), pid)
		logger.WithContext(ctx).Infof("[%s]%s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
