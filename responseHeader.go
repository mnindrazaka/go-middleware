package main

import "net/http"

type responseHeader struct {
	handler     http.Handler
	headerName  string
	headerValue string
}

type ResponseHeader interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func NewResponseHeader(handler http.Handler, headerName string, headerValue string) ResponseHeader {
	return responseHeader{handler, headerName, headerValue}
}

func (responseHeader responseHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(responseHeader.headerName, responseHeader.headerValue)
	responseHeader.handler.ServeHTTP(w, r)
}
