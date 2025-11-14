package main

import (
	"net/http"
	"os"
	"os/signal"
	"shorten-url/app"
	"shorten-url/package/logger"
)

func main() {
	runApplication()
}

func runApplication() {
	app := app.NewApplication()
	keepAliveSignal := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := app.Stop(); err != nil {
			logger.Error(err, "fail to shutdown")
		}
		close(keepAliveSignal)
	}()
	if err := app.Start(); err != nil && err != http.ErrServerClosed {
		logger.Error(err, "fail to start app")
	}
	<-keepAliveSignal
	logger.Info("Application stopped")
}
