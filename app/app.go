package app

import (
	"context"
	"fmt"
	"net/http"
	"shorten-url/internal/adapter/handler"
	"shorten-url/internal/adapter/repository"
	"shorten-url/internal/router"
	"shorten-url/internal/service"
	"shorten-url/package/database/mysql"
	"shorten-url/package/httputils"
	"shorten-url/package/logger"
)

type Application struct {
	server *http.Server
	stop   func() error
}

func NewApplication() *Application {
	cfg := GetConfig()
	logger.InitZeroLogger(cfg.LogLevel)
	logger.Debugf("Config:\n%s", logger.JSONFormat(cfg))
	handler, stop := setupServerHandler(cfg)
	server := &http.Server{
		Handler: handler,
		Addr: fmt.Sprintf(
			"%s:%s",
			cfg.ServiceConfig.Host,
			cfg.ServiceConfig.Port,
		),
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}
	return &Application{
		server: server,
		stop:   stop,
	}
}

func setupServerHandler(cfg Config) (http.Handler, func() error) {
	db, err := mysql.NewClient(cfg.MySQLConfig)
	if err != nil {
		logger.Fatal(err)
	}
	stopFunc := func() error {
		return db.Close()
	}
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)
	router := router.NewRouter(handler)
	return httputils.GetCORSMiddleWare(cfg.CORSConfig)(router), stopFunc

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
