package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router *gin.Engine

func NewRouter(handler *Handler) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	deliveries := router.Group("/deliveries")
	{
		deliveries.POST("/add_delivery", handler.CreateDelivery)
		deliveries.GET("/list", handler.GetDeliveries)
		deliveries.GET("/route_distance", handler.GetRouteDistance)
		deliveries.GET("/:id", handler.GetDelivery)
	}

	return router
}
