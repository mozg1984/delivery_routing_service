package app

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
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
