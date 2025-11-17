package router

import (
	"net/http"
	"shorten-url/internal/port"
	"shorten-url/package/httputils"
	"shorten-url/package/middleware"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(handler port.Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router.Use(
		middleware.LogMiddleware,
		middleware.RecoverMiddleware,
	)
	router.HandleFunc("/health", httputils.CheckHealth).Methods(http.MethodGet)
	router.HandleFunc("/encode", handler.Encode).Methods(http.MethodPost)
	router.HandleFunc("/decode", handler.Decode).Methods(http.MethodPost)
	return router
}
