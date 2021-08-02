package app

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
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

func (h *Handler) CreateDelivery(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Create delivery")
}

func (h *Handler) GetDeliveries(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Get list of deliveries ordered by created time")
}

func (h *Handler) GetDelivery(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Get delivery")
}

func (h *Handler) GetRouteDistance(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Calculate route distance and return it")
}
