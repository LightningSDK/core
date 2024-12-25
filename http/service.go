package http

import (
	"github.com/lightningsdk/core/model"
	"log"
	"net/http"
	"sync"
)

func GetCommand() model.Command {
	return model.Command{
		Function: nil,
		Help:     "commands for the http server",
		SubCommands: map[string]model.Command{
			"start": model.Command{
				Function: func(a model.App) error {
					s := NewService("8080")
					return s.Start(a)
				},
				Help: "starts the http server",
			},
		},
	}
}

type Service struct {
	port   string
	router *Router
	mutex  sync.RWMutex
}

// NewService creates a new Service with the provided port
func NewService(port string) *Service {
	return &Service{
		port:  port,
		mutex: sync.RWMutex{},
	}
}

// AddRouter registers a new Router for a specific domain
func (s *Service) SetRouter(router *Router) {
	s.router = router
}

// Start begins listening on the specified port and routes requests to the appropriate router
func (s *Service) Start(a model.App) error {
	s.router = NewRouter()
	s.router.LoadRoutes(a)
	log.Printf("Service is listening on port %s...\n", s.port)
	return http.ListenAndServe(":"+s.port, s.router)
}
