package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

// constant for server port
const SERVER_PORT string = ":8080"

// logger initialization
var logger = log.New(os.Stdout, "go-web-server", log.LstdFlags|log.Lshortfile)

// middleware for logging request processing time
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next(w, r)
		logger.Printf("Request for %s processed in %s \n", r.URL.Host+r.URL.Path, time.Since(startTime))
	}
}
