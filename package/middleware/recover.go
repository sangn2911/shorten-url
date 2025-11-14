package middleware

import (
	"net/http"
	"runtime/debug"
	"shorten-url/package/logger"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if value := recover(); value != nil {
				logger.Errorf(nil,
					"panic \"%v\" with stack trace:\n%s",
					value,
					debug.Stack(),
				)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
