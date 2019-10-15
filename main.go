package main

import (
	"fmt"
	"github.com/MateuszPu/go-web-app-starter/config"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", config.LoggerHandler(hello()).ServeHTTP)


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

func hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := r.URL.Query().Get("value")
		fmt.Fprintf(w, "Hello %s!", value)
	})
}