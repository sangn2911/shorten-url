package handler

import (
	"net/http"
	"shorten-url/internal/adapter/handler/dto"
	"shorten-url/internal/port"
	"shorten-url/internal/service"
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
	logger.Info("Handler.Encode")
	var request dto.ConvertUrlRequest
	if err := httputils.DecodeRequest(r, &request); err != nil {
		httputils.JSON400(w, r, err.Error())
		return
	}
	shortenUrl, err := h.service.Encode(r.Context(), request.Url)
	if err != nil {
		switch err {
		case service.ErrUrlEmpty:
			httputils.JSON400(w, r, err.Error())
			return
		default:
			httputils.JSON500(w, r, err.Error())
		}
		return
	}
	httputils.JSON200(w, r, dto.EncodeUrlResponse{
		ShortenUrl: shortenUrl,
	})
}

func (h *handler) Decode(w http.ResponseWriter, r *http.Request) {
	logger := logger.WithContext(r.Context())
	logger.Info("Handler.Decode")
	var request dto.ConvertUrlRequest
	if err := httputils.DecodeRequest(r, &request); err != nil {
		httputils.JSON400(w, r, err.Error())
		return
	}
	originalUrl, err := h.service.Decode(r.Context(), request.Url)
	if err != nil {
		switch err {
		case service.ErrShortenUrlNotFound:
			httputils.JSON404(w, r, err.Error())
		case service.ErrUrlEmpty:
			httputils.JSON400(w, r, err.Error())
			return
		default:
			httputils.JSON500(w, r, err.Error())
		}
		return
	}
	httputils.JSON200(w, r, dto.DecodeUrlResponse{
		OriginalUrl: originalUrl,
	})
}
