package app

import "github.com/julienschmidt/httprouter"

func Routes() *httprouter.Router {
	router := httprouter.New()

	{
		router.POST("/deliveries/add_delivery", CreateDelivery)
		router.GET("/deliveries/list", GetDeliveries)
		router.GET("/deliveries/", GetDelivery)
		router.GET("/deliveries/route_distance", GetRouteDistance)
	}

	return router
}
