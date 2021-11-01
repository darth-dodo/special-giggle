package main

import (
	"fmt"
	"github.com/darth-dodo/special-giggle/books-api/app/router"
	"log"
	"net/http"

	"github.com/darth-dodo/special-giggle/books-api/config"
)

func main() {

	appConf := config.AppConfig()
	appRouter := router.New()

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	log.Printf("Starting server %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello World!")
	if err != nil {
		return
	}
}
