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
		//SubCommands: map[string]model.Command{
		//	"start": model.Command{
		//		Function: func(a model.App) {
		//
		//		},
		//	},
		//},
	}
}

// Router interface is expected to be defined in ./router.go
type HttpService interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Service struct {
	port         string
	domainRouter map[string]Router
	mutex        sync.RWMutex
}

// NewService creates a new Service with the provided port
func NewService(port string) *Service {
	return &Service{
		port:         port,
		domainRouter: make(map[string]Router),
		mutex:        sync.RWMutex{},
	}
}

// AddRouter registers a new Router for a specific domain
func (s *Service) AddRouter(domain string, router Router) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.domainRouter[domain] = router
}

// Start begins listening on the specified port and routes requests to the appropriate router
func (s *Service) Start() error {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		domain := r.Host // Extract the domain from the request

		s.mutex.RLock()
		router, exists := s.domainRouter[domain]
		s.mutex.RUnlock()

		if exists {
			router.ServeHTTP(w, r)
		} else {
			http.Error(w, "Domain not found", http.StatusNotFound)
		}
	})

	log.Printf("Service is listening on port %s...\n", s.port)
	return http.ListenAndServe(":"+s.port, handler)
}

//func main() {
//	// Create a new service instance
//	service := NewService("8080")
//
//	// Example routers (implementations must exist in ./router.go)
//	exampleRouter := &ExampleRouter{}
//	anotherRouter := &AnotherRouter{}
//
//	// Add routers for specific domains
//	service.AddRouter("example.com", exampleRouter)
//	service.AddRouter("anotherdomain.com", anotherRouter)
//
//	// Start the service
//	if err := service.Start(); err != nil {
//		log.Fatal(err)
//	}
//}
//
//// ExampleRouter is a placeholder for an actual Router implementation
//type ExampleRouter struct{}
//
//func (e *ExampleRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Hello from ExampleRouter"))
//}
//
//// AnotherRouter is another placeholder for a Router implementation
//type AnotherRouter struct{}
//
//func (a *AnotherRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Welcome to AnotherRouter"))
//}
