package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mozg1984/delivery_routing_service/pkg/api"
)

type Handler struct {
	deliveryService api.DeliveryService
}

func NewHandler(deliveryService api.DeliveryService) *Handler {
	return &Handler{
		deliveryService: deliveryService,
	}
}

func (h *Handler) CreateDelivery(c *gin.Context) {
	newDeliveryRequest := api.NewDeliveryRequest{}

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&newDeliveryRequest)
	if err != nil {
		log.Printf("An error occurred while decoding in the creating delivery: '%v'", err)
		c.String(http.StatusBadRequest, "Invalid params")
		return
	}

	err = h.deliveryService.New(newDeliveryRequest)
	if err != nil {
		log.Printf("An error occurred while calling api client in the creating delivery: '%v'", err)
		c.String(http.StatusBadRequest, "Invalid params")
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) GetDeliveries(c *gin.Context) {
	deliveries, err := h.deliveryService.GetAll()
	if err != nil {
		log.Printf("An error occurred while calling api client in getting delivery: '%v'", err)
		c.String(http.StatusBadRequest, "Something went wrong, please try later")
		return
	}

	c.JSON(http.StatusOK, deliveries)
}

func (h *Handler) GetDelivery(c *gin.Context) {
	id := c.Param("id")

	deliveryId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Printf("An error occurred while converting param in getting delivery: '%v'", err)
		c.JSON(http.StatusOK, nil)
		return
	}

	delivery, err := h.deliveryService.FingByID(api.DeliveryID(deliveryId))
	if err != nil {
		log.Printf("An error occurred while calling api client in getting delivery: '%v'", err)
		c.JSON(http.StatusOK, nil)
		return
	}

	c.JSON(http.StatusOK, delivery)
}

func (h *Handler) GetRouteDistance(c *gin.Context) {
	routeDistance, err := h.deliveryService.CalculateRouteDistance()
	if err != nil {
		log.Printf("An error occurred while calling api client in getting route distance: '%v'", err)
		c.String(http.StatusOK, "-1")
		return
	}

	c.String(http.StatusOK, strconv.FormatFloat(routeDistance, 'E', -1, 64))
}
