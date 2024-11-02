package api

import (
	"log"
	"net/http"

	"github.com/KengoWada/gorouting/services/user"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Run() error {
	userHandler := user.NewHandler()
	userRouter := userHandler.RegisterRoutes()

	router := http.NewServeMux()
	router.Handle("/api/v1/users/", http.StripPrefix("/api/v1/users", userRouter))

	log.Println("Starting server on port", s.addr)
	return http.ListenAndServe(s.addr, router)
}
