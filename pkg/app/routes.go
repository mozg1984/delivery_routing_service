package app

import "github.com/julienschmidt/httprouter"

type Router *httprouter.Router

func NewRouter(handler *Handler) *httprouter.Router {
	router := httprouter.New()

	{
		router.POST("/deliveries/add_delivery", handler.CreateDelivery)
		router.GET("/deliveries/list", handler.GetDeliveries)
		router.GET("/deliveries/", handler.GetDelivery)
		router.GET("/deliveries/route_distance", handler.GetRouteDistance)
	}

	return router
}
