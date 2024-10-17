package handlers

import (
	"net/http"
)

type Route struct {
	method  string
	path    string
	handler func(http.ResponseWriter, *http.Request)
}
