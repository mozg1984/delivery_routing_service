package api

import (
	"errors"
	"reflect"
	"testing"
)

type mockDeliveryRepo struct{}

func (m *mockDeliveryRepo) CreateDelivery(request NewDeliveryRequest) error {
	return nil
}

func (m *mockDeliveryRepo) GetDeliveries() (*[]Delivery, error) {
	// TODO: implement it
	var delivery Delivery
	return &[]Delivery{delivery}, nil
}

func (m *mockDeliveryRepo) GetDelivery(deliveryId DeliveryID) (Delivery, error) {
	// TODO: implement it
	var delivery Delivery
	return delivery, nil
}

func Test_deliveryService_New(t *testing.T) {
	mockRepo := mockDeliveryRepo{}
	mockDeliveryServis := NewDeliveryService(&mockRepo)

	tests := []struct {
		name    string
		request NewDeliveryRequest
		wantErr error
	}{
		{name: "", request: NewDeliveryRequest{
			ID:           DeliveryID(123456),
			Pickup:       Coordinates{Latitude: 123.000, Longitude: 234.000},
			Dropoff:      Coordinates{Latitude: 122.999, Longitude: 234.999},
			CreationTime: "2021-07-21T19:37:06+00:00",
		}, wantErr: nil},
		{name: "", request: NewDeliveryRequest{
			ID:           DeliveryID(0),
			Pickup:       Coordinates{Latitude: 123.000, Longitude: 234.000},
			Dropoff:      Coordinates{Latitude: 122.999, Longitude: 234.999},
			CreationTime: "2021-07-21T19:37:06+00:00",
		}, wantErr: errors.New("delivery service - ID required")},
		{name: "", request: NewDeliveryRequest{
			ID:           DeliveryID(123456),
			Pickup:       Coordinates{Latitude: 123.000, Longitude: 234.000},
			Dropoff:      Coordinates{Latitude: 122.999, Longitude: 234.999},
			CreationTime: "",
		}, wantErr: errors.New("delivery service - Creation time required")},
		{name: "", request: NewDeliveryRequest{
			ID:           DeliveryID(123456),
			Pickup:       Coordinates{Latitude: 123.000, Longitude: 234.000},
			Dropoff:      Coordinates{Latitude: 122.999, Longitude: 234.999},
			CreationTime: "2021-07-21 19:37:06+00:00",
		}, wantErr: errors.New("delivery service - Creation time should have RFC3339 format")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := mockDeliveryServis.New(test.request)

			if !reflect.DeepEqual(err, test.wantErr) {
				t.Errorf("test: %v failed. got: %v, wanted: %v", test.name, err, test.wantErr)
			}
		})
	}
}
