package main

import (
	"log"
	"net/http"
	"os"
)

type Service struct {
	Logger *log.Logger
}

func main() {
	// Create a new service
	s := &Service{
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}

	// Run the service
	if err := s.Run(); err != nil {
		s.Logger.Fatal(err)
	}
}

func (s *Service) Run() error {

	return nil
}

func RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	return r.ServeHTTP
}
