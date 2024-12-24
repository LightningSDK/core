package http

import (
	"github.com/lightningsdk/core/model"
	"net/http"
	"sort"
	"sync"
)

// Router represents the routing service.
type Router struct {
	routes []model.Route
	lock   sync.RWMutex
}

// NewRouter initializes and returns a new Router instance.
func NewRouter() *Router {
	return &Router{
		routes: make([]model.Route, 0),
	}
}

// AddRoutes allows adding a list of routes to the service.
func (s *Router) AddRoutes(routes []model.Route) {
	s.lock.Lock()
	defer s.lock.Unlock()
	// Append the provided routes to the service's route list.
	s.routes = append(s.routes, routes...)

	// Sort the routes by weight.
	sort.Slice(s.routes, func(i, j int) bool {
		return s.routes[i].Weight < s.routes[j].Weight
	})
}

// ServeHTTP implements http.Handler, allowing the service to be used as a handler.
func (s *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	// Find a matching route handler.
	for _, route := range s.routes {
		if r.URL.Path == route.Path {
			// Execute the handler for the matching route.
			route.Handler(w, r)
			return
		}
	}
	// If no match was found, send a 404 response.
	http.NotFound(w, r)
}

// Example usage
func example() {
	// Create a new service.
	service := NewRouter()

	// Define some routes.
	routes := []model.Route{
		{
			Path:   "/hello",
			Weight: 1,
			Handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Hello, world!"))
			},
		},
		{
			Path:   "/goodbye",
			Weight: 2,
			Handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Goodbye, world!"))
			},
		},
	}

	// Add the defined routes to the service.
	service.AddRoutes(routes)

	// Use the service with an HTTP server.
	http.ListenAndServe(":8080", service)
}
