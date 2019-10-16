package config

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func init() {
	toFile, err := os.OpenFile("application.log", os.O_WRONLY | os.O_CREATE, 0755)
	if err != nil {
	}
	log.SetOutput(toFile)
}

func LoggerHandler(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()
		inner.ServeHTTP(response, request)
		log.Infof("Request method: [%s], request url: [%s], latency: [%.2f]", request.Method,  request.URL.String(), time.Since(start).Milliseconds())
	})
}
