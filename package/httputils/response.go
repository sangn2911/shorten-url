package httputils

import (
	"encoding/json"
	"net/http"
	"shorten-url/package/logger"
)

type Response struct {
	StatusCode int    `json:"-"`
	Data       any    `json:"-"`
	Message    string `json:"message"`
}

func JSON200(w http.ResponseWriter, r *http.Request, data any) {
	JSON(w, r, Response{
		Data:       data,
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
	})
}

func JSON400(w http.ResponseWriter, r *http.Request, message string) {
	JSON(w, r, Response{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	})
}

func JSON500(w http.ResponseWriter, r *http.Request, message string) {
	JSON(w, r, Response{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	})
}

func JSON404(w http.ResponseWriter, r *http.Request, message string) {
	JSON(w, r, Response{
		StatusCode: http.StatusNotFound,
		Message:    message,
	})
}

func JSON(w http.ResponseWriter, r *http.Request, res Response) {
	logger := logger.WithContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	var response any = res
	if res.Data != nil {
		response = res.Data
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
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error(err, "fail to send response")
	}
}
