package app

import "github.com/julienschmidt/httprouter"

func (s *Server) Routes() *httprouter.Router {
	router := httprouter.New()

	{
		router.POST("/deliveries/add_delivery", s.handler.CreateDelivery)
		router.GET("/deliveries/list", s.handler.GetDeliveries)
		router.GET("/deliveries/", s.handler.GetDelivery)
		router.GET("/deliveries/route_distance", s.handler.GetRouteDistance)
	}

	return router
}
