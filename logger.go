package main

import (
	"log"
	"net/http"
	"time"
)

type logger struct {
	handler http.Handler
}

type Logger interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func NewLogger(handler http.Handler) Logger {
	return logger{handler}
}

func (logger logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	logger.handler.ServeHTTP(w, r)
}
