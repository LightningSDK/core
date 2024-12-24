package model

import "net/http"

// Route represents a structure for storing route configuration.
type Route struct {
	Path    string
	Weight  int
	Handler http.HandlerFunc
}
