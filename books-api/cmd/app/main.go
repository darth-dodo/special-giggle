package main

import (
	"fmt"
	"github.com/darth-dodo/special-giggle/books-api/app/app"
	"github.com/darth-dodo/special-giggle/books-api/app/router"
	"net/http"

	"github.com/darth-dodo/special-giggle/books-api/config"
	lr "github.com/darth-dodo/special-giggle/books-api/util/logger"
)

func main() {

	appConf := config.AppConfig()
	logger := lr.New(appConf.Debug)
	application := app.New(logger)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)
	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello World!")
	if err != nil {
		return
	}
}
