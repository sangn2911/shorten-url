package httputils

import (
	"encoding/json"
	"net/http"
	"shorten-url/package/logger"
)

type response struct {
	StatusCode int    `json:"-"`
	Data       any    `json:"data"`
	Message    string `json:"message"`
}

func JSON200(w http.ResponseWriter, r *http.Request, data any) {
	writeResponse(w, r, response{
		Data:       data,
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
	})
}

func JSON400(w http.ResponseWriter, r *http.Request, data any, message string) {
	writeResponse(w, r, response{
		Data:       data,
		StatusCode: http.StatusBadRequest,
		Message:    message,
	})
}

func JSON500(w http.ResponseWriter, r *http.Request, message string) {
	writeResponse(w, r, response{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	})
}

func writeResponse(w http.ResponseWriter, r *http.Request, res response) {
	logger := logger.WithContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	var data any = res
	if res.Data != nil {
		data = res.Data
	}
	if len(res.Message) > 0 {
		logger.Infof(
			"[%s]%s status=%d message=%s",
			r.Method,
			r.URL.Path,
			res.StatusCode,
			res.Message,
		)
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error(err, "fail to send response")
	}
}
