package main

import (
	"fmt"
	"net/http"
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

	app.logger.Info(
		"starting api server",
		zap.String("addr", srv.Addr),
		zap.String("env", app.config.env))
	return srv.ListenAndServe()

}
