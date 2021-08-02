package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mozg1984/delivery_routing_service/pkg/api"
)

type Server struct {
	router          *httprouter.Router
	handler         *Handler
	deliveryService api.DeliveryService
}

func NewServer(router *httprouter.Router, handler *Handler, deliveryService api.DeliveryService) *Server {
	return &Server{
		router:          router,
		handler:         handler,
		deliveryService: deliveryService,
	}
}

func (s *Server) Run() error {
	return nil
}
