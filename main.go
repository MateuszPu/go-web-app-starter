package main

import (
	"github.com/MateuszPu/go-web-app-starter/config"
	"github.com/MateuszPu/go-web-app-starter/user"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", config.LoggerHandler(user.Get()).ServeHTTP)


	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
