package httputils

import (
	"context"
	"encoding/json"
	"net/http"
	"shorten-url/package/logger"
)

type response struct {
	StatusCode int    `json:"-"`
	Data       any    `json:"data"`
	Message    string `json:"message"`
}

func JSON200(ctx context.Context, w http.ResponseWriter, data any) {
	writeResponse(w, response{
		Data:       data,
		StatusCode: http.StatusOK,
	})
}

func JSON400(ctx context.Context, w http.ResponseWriter, data any, message string) {
	writeResponse(w, response{
		Data:       data,
		StatusCode: http.StatusBadRequest,
		Message:    message,
	})
}

func JSON500(ctx context.Context, w http.ResponseWriter, message string) {
	writeResponse(w, response{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	})
}

func writeResponse(w http.ResponseWriter, res response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	var data any = res
	if res.Data != nil {
		data = res.Data
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error(err, "fail to send response")
	}
}
