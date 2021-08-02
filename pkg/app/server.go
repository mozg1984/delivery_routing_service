package app

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mozg1984/delivery_routing_service/pkg/config"
)

type Server struct {
	router *httprouter.Router
}

func NewServer(router *httprouter.Router) *Server {
	return &Server{
		router: router,
	}
}

func (s *Server) Run(config *config.Server) error {
	var addr string = config.Host + ":" + config.Port

	log.Print("Starting server")

	err := http.ListenAndServe(addr, s.router)
	if err != nil {
		log.Printf("An error occurred while starting the server: %v", err)
		return err
	}

	return nil
}
