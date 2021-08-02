package app

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mozg1984/delivery_routing_service/pkg/api"
	"github.com/mozg1984/delivery_routing_service/pkg/config"
)

type Server struct {
	router          *httprouter.Router
	deliveryService api.DeliveryService
}

func NewServer(router *httprouter.Router, deliveryService api.DeliveryService) *Server {
	return &Server{
		router:          router,
		deliveryService: deliveryService,
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
