package middleware

import (
	"net/http"
	"shorten-url/package/logger"

	"github.com/google/uuid"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pid := uuid.NewString()
		logger.Infof("pid=%s [%s]%s", pid, r.Method, r.URL.Path)
		ctx := logger.SetPID(r.Context(), pid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
