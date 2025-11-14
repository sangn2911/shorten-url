package handler

import (
	"net/http"
	"shorten-url/internal/model"
	"shorten-url/internal/port"
	"shorten-url/package/httputils"
	"shorten-url/package/logger"
)

type handler struct {
	service port.Service
}

func NewHandler(service port.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Encode(w http.ResponseWriter, r *http.Request) {
	logger := logger.WithContext(r.Context())
	logger.Info("Parse request...")
	var request model.ConvertUrlRequest
	if err := httputils.DecodeRequest(r, &request); err != nil {
		httputils.JSON400(r.Context(), w, nil, err.Error())
		return
	}
	shortenUrl, err := h.service.Encode(r.Context(), request.Url)
	if err != nil {
		httputils.JSON500(r.Context(), w, err.Error())
		return
	}
	httputils.JSON200(r.Context(), w, model.EncodeUrlResponse{
		ShortenUrl: shortenUrl,
	})
}

func (h *handler) Decode(w http.ResponseWriter, r *http.Request) {
	logger := logger.WithContext(r.Context())
	logger.Info("Parse request...")
	var request model.ConvertUrlRequest
	if err := httputils.DecodeRequest(r, &request); err != nil {
		httputils.JSON400(r.Context(), w, nil, err.Error())
		return
	}
	originalUrl, err := h.service.Decode(r.Context(), request.Url)
	if err != nil {
		httputils.JSON500(r.Context(), w, err.Error())
		return
	}
	httputils.JSON200(r.Context(), w, model.DecodeUrlResponse{
		OriginalUrl: originalUrl,
	})
}
