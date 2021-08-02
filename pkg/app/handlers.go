package app

import (
	"fmt"

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
	fmt.Fprint(c.Writer, "Create delivery")

	//newDeliveryRequest := api.NewDeliveryRequest{}

	//decoder := json.NewDecoder(r.Body)
	//err := decoder.Decode(&newDeliveryRequest)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	//err = h.deliveryService.New(newDeliveryRequest)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	//w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetDeliveries(c *gin.Context) {
	fmt.Fprint(c.Writer, "Get list of deliveries ordered by created time")

	//deliveries := h.deliveryService.GetAll()
}

func (h *Handler) GetDelivery(c *gin.Context) {
	fmt.Fprint(c.Writer, "Get delivery")

	//deliveryId := params.ByName("id")
	////h.deliveryService.FingByID()

	//fmt.Fprint(w, deliveryId)
}

func (h *Handler) GetRouteDistance(c *gin.Context) {
	fmt.Fprint(c.Writer, "Get route distance")
	//routeDistance := h.deliveryService.CalculateRouteDistance()
}
