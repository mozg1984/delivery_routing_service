package api

type Coordinates struct {
	Latitude  float32 `json:"lan"`
	Longitude float32 `json:"lng"`
}

type DeliveryID uint

type Delivery struct {
	ID           DeliveryID  `json:"id"`
	Pickup       Coordinates `json:"pickup"`
	Dropoff      Coordinates `json:"dropoff"`
	CreationTime string      `json:"creation_time"`
}

type NewDeliveryRequest struct {
	ID           DeliveryID  `json:"id"`
	Pickup       Coordinates `json:"pickup"`
	Dropoff      Coordinates `json:"dropoff"`
	CreationTime string      `json:"creation_time"`
}
