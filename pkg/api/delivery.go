package api

import (
	"errors"
	"regexp"

	geo "github.com/mozg1984/delivery_routing_service/pkg/calculator"
)

const RFC3339_PATTERN string = `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}`

type DeliveryService interface {
	New(delivery NewDeliveryRequest) error
	GetAll() (*[]Delivery, error)
	FingByID(DeliveryID) (Delivery, error)
	CalculateRouteDistance() (float64, error)
}

type deliveryService struct {
	storage DeliveryRepository
}

type DeliveryRepository interface {
	CreateDelivery(request NewDeliveryRequest) error
	GetDeliveries(flushdb bool) (*[]Delivery, error)
	GetDelivery(deliveryId DeliveryID) (Delivery, error)
}

func NewDeliveryService(userRepo DeliveryRepository) DeliveryService {
	return &deliveryService{storage: userRepo}
}

func (d *deliveryService) New(delivery NewDeliveryRequest) error {
	err := d.validate(delivery)
	if err != nil {
		return err
	}

	err = d.storage.CreateDelivery(delivery)
	if err != nil {
		return err
	}

	return nil
}

func (d *deliveryService) validate(delivery NewDeliveryRequest) error {
	if delivery.ID == 0 {
		return errors.New("delivery service - ID required")
	}

	if delivery.CreationTime == "" {
		return errors.New("delivery service - Creation time required")
	}

	matched, err := regexp.MatchString(RFC3339_PATTERN, delivery.CreationTime)
	if err != nil {
		return err
	}

	if !matched {
		return errors.New("delivery service - Creation time should have RFC3339 format")
	}

	return nil
}

func (d *deliveryService) GetAll() (*[]Delivery, error) {
	return d.storage.GetDeliveries(true)
}

func (d *deliveryService) FingByID(deliveryId DeliveryID) (Delivery, error) {
	return d.storage.GetDelivery(deliveryId)
}

func (d *deliveryService) CalculateRouteDistance() (float64, error) {
	routeDistance := float64(0)

	deliveries, err := d.storage.GetDeliveries(false)
	if err != nil {
		return routeDistance, nil
	}

	for _, delivery := range *deliveries {
		routeDistance += geo.CalculateDistance(
			delivery.Pickup.Latitude,
			delivery.Pickup.Longitude,
			delivery.Dropoff.Latitude,
			delivery.Dropoff.Longitude,
		)
	}

	return routeDistance, nil
}
