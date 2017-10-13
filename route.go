package main

import (
	"net/http"
)

const (
	allowedHttpMethods = []string{"GET", "POST", "HEAD", "PATCH", "PUT", "DELETE", "TRACE", "CONNECT", "OPTIONS"}
)

type Route struct {
	methods []string

	path string

	handler http.Handler
}
