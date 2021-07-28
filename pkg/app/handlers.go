package app

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateDelivery(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Create delivery")
}

func GetDeliveries(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Get list of deliveries ordered by created time")
}

func GetDelivery(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Get delivery")
}

func GetRouteDistance(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Calculate route distance and return it")
}
