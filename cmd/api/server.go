package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func (app *application) serve() error {
	errorLogger, _ := zap.NewStdLogAt(app.logger, zap.ErrorLevel)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		ErrorLog:     errorLogger,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		app.logger.Info("caught signal", zap.String("signal", s.String()))

		os.Exit(0)
	}()

	app.logger.Info(
		"starting api server",
		zap.String("addr", srv.Addr),
		zap.String("env", app.config.env))
	return srv.ListenAndServe()

}
