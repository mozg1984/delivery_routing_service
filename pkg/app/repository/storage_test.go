package repository

import (
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/mozg1984/delivery_routing_service/pkg/api"
	"github.com/rafaeljusto/redigomock"
)

func Test_storage_CreateDelivery(t *testing.T) {
	redisConnectionMock, testStorage := mockRedisStorage()
	redisConnectionMock.Clear()

	request := api.NewDeliveryRequest{
		ID: api.DeliveryID(123456),
		Pickup: api.Coordinates{
			Latitude:  123.456,
			Longitude: 234.567,
		},
		Dropoff: api.Coordinates{
			Latitude:  345.678,
			Longitude: 456.789,
		},
		CreationTime: "2021-07-21T19:37:06+00:00",
	}

	setNXCmd := redisConnectionMock.Command(
		"setnx",
		"delivery:123456",
		`{"id":123456,"pickup":{"lat":123.456,"lng":234.567},"dropoff":{"lat":345.678,"lng":456.789},"creation_time":"2021-07-21T19:37:06+00:00"}`,
	).Expect("OK!")

	zAddCmd := redisConnectionMock.Command("zadd", "ids", "nx", int64(1626896226), "delivery:123456").Expect("OK!")

	err := testStorage.CreateDelivery(request)
	if err != nil {
		t.Errorf("Delivery creation is failed: %s", err)
	}

	if redisConnectionMock.Stats(setNXCmd) != 1 {
		t.Fatal("Command SETNX was not called!")
	}

	if redisConnectionMock.Stats(zAddCmd) != 1 {
		t.Fatal("Command ZADD was not called!")
	}
}

func Test_storage_GetDeliveries(t *testing.T) {
	redisConnectionMock, testStorage := mockRedisStorage()
	redisConnectionMock.Clear()

	values := []interface{}{}
	values = append(values, interface{}([]byte("delivery:123456")))
	values = append(values, interface{}([]byte("delivery:123457")))
	values = append(values, interface{}([]byte("delivery:123458")))

	fmt.Println(values)

	commands := []*redigomock.Cmd{
		redisConnectionMock.GenericCommand("zrange").Expect(values),
		redisConnectionMock.Command("get", "delivery:123456").Expect(
			`{"id":123456,"pickup":{"lat":123.456,"lng":234.567},"dropoff":{"lat":345.678,"lng":456.789},"creation_time":"2021-07-21T19:37:06+00:00"}`,
		),
		redisConnectionMock.Command("get", "delivery:123457").Expect(
			`{"id":123457,"pickup":{"lat":789.012,"lng":890.123},"dropoff":{"lat":890.123,"lng":901.234},"creation_time":"2021-07-21T19:38:06+00:00"}`,
		),
		redisConnectionMock.Command("get", "delivery:123458").Expect(
			`{"id":123458,"pickup":{"lat":901.234,"lng":321.654},"dropoff":{"lat":275.778,"lng":111.245},"creation_time":"2021-07-21T19:39:06+00:00"}`,
		),
		redisConnectionMock.Command("flushdb", "async"),
	}

	deliveries, err := testStorage.GetDeliveries(true)
	if err != nil {
		t.Errorf("Getting deliveries is failed: %s", err)
	}

	if len(*deliveries) != 3 {
		t.Fatal("Not all deliveries have been retrieved!")
	}

	for _, cmd := range commands {
		if redisConnectionMock.Stats(cmd) != 1 {
			t.Fatal("Command was not called!")
		}
	}
}

func Test_storage_GetDelivery(t *testing.T) {
	redisConnectionMock, testStorage := mockRedisStorage()
	redisConnectionMock.Clear()

	deliveryId := api.DeliveryID(123456)

	cmd := redisConnectionMock.Command("get", "delivery:123456").Expect(
		`{"id":123456,"pickup":{"lat":123.456,"lng":234.567},"dropoff":{"lat":345.678,"lng":456.789},"creation_time":"2021-07-21T19:37:06+00:00"}`,
	)

	delivery, err := testStorage.GetDelivery(deliveryId)
	if err != nil {
		t.Errorf("Getting a delivery is failed: %s", err)
	}

	if redisConnectionMock.Stats(cmd) != 1 {
		t.Fatal("Command GET was not called!")
	}

	if delivery.ID != deliveryId {
		t.Errorf("Invalid delivery ID. Expected '123456' and got '%d'", delivery.ID)
	}
}

func mockRedisStorage() (*redigomock.Conn, *storage) {
	redisConnectionMock := redigomock.NewConn()

	testStorage := &storage{
		pool: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redisConnectionMock, nil
			},
		},
	}

	return redisConnectionMock, testStorage
}
