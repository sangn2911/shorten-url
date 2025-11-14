package app

import (
	"context"
	"fmt"
	"net/http"
	"shorten-url/package/httputils"
	"shorten-url/package/logger"
	"shorten-url/package/middleware"

	"github.com/gorilla/mux"
)

type Application struct {
	server *http.Server
	stop   func() error
}

func NewApplication() *Application {
	cfg := GetConfig()
	logger.InitZeroLogger(cfg.LogLevel)
	logger.Debugf("Config:\n%s", logger.JSONFormat(cfg))
	server := &http.Server{
		Handler:      setupServerHandler(cfg),
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}
	return &Application{
		server: server,
		stop: func() error {
			return nil
		},
	}
}

func setupServerHandler(cfg Config) http.Handler {
	_ = cfg
	router := mux.NewRouter().StrictSlash(true)
	router.Use(
		middleware.LogMiddleware,
		middleware.RecoverMiddleware,
	)
	router.HandleFunc("/health", httputils.CheckHealth).Methods(http.MethodGet)
	return httputils.GetCORSMiddleWare(cfg.CORSConfig)(router)
}

func (a *Application) Start() error {
	logger.Infof("Start application at %s", a.server.Addr)
	return a.server.ListenAndServe()
}

func (a *Application) Stop() error {
	logger.Info("Shutdown application...")
	if err := a.stop(); err != nil {
		logger.Errorf(err, "fail to stop connections")
	}
	return a.server.Shutdown(context.Background())
}
