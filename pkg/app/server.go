package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mozg1984/delivery_routing_service/pkg/config"
)

type Server struct {
	router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
	return &Server{
		router: router,
	}
}

func (s *Server) Run(config *config.Server) error {
	var addr string = config.Host + ":" + config.Port

	log.Print("Starting server")

	err := http.ListenAndServe(addr, s)
	if err != nil {
		log.Printf("An error occurred while starting the server: %v", err)
		return err
	}

	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	s.router.ServeHTTP(w, r)
}
