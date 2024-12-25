package http

import (
	"github.com/lightningsdk/core/model"
	"net/http"
	"sort"
)

// Router represents the routing service.
type Router struct {
	routes []model.Route
}

// NewRouter initializes and returns a new Router instance.
func NewRouter() *Router {
	return &Router{
		routes: make([]model.Route, 0),
	}
}

// AddRoutes allows adding a list of routes to the service.
func (s *Router) LoadRoutes(a model.App) {
	s.routes = []model.Route{}

	// Iterate over all plugins and collect their routes.
	for _, plugin := range a.GetPlugins() {
		pluginRoutes := plugin.GetRoutes()
		s.routes = append(s.routes, pluginRoutes...)
	}

	// Sort the routes by weight.
	sort.Slice(s.routes, func(i, j int) bool {
		return s.routes[i].Weight < s.routes[j].Weight
	})
}

// ServeHTTP implements http.Handler, allowing the service to be used as a handler.
func (s *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
