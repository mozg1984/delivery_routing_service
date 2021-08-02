package repository

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/mozg1984/delivery_routing_service/pkg/api"
	"github.com/mozg1984/delivery_routing_service/pkg/config"
)

type Storage interface {
	CreateDelivery(request api.NewDeliveryRequest) error
	GetDeliveries(flushdb bool) (*[]api.Delivery, error)
	GetDelivery(deliveryId api.DeliveryID) (api.Delivery, error)
	GetRawValues(flushdb bool) (*[]string, error)
	GetRawValue(key string) (string, error)
}

type storage struct {
	pool *redis.Pool
}

type Store struct{}

func NewStorage(c *config.Redis) Storage {
	var addr string = c.Host + ":" + c.Port

	pool := &redis.Pool{
		MaxIdle:     c.MaxIdle,
		MaxActive:   c.MaxActive,
		IdleTimeout: time.Duration(c.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(c.Type, addr)
			if err != nil {
				return nil, err
			}

			return conn, nil
		},
	}

	return &storage{pool: pool}
}

func (s *storage) CreateDelivery(request api.NewDeliveryRequest) error {
	conn := s.pool.Get()
	defer conn.Close()

	var (
		deliveryId string = strconv.FormatUint(uint64(request.ID), 10)
		key        string = "delivery:" + deliveryId
	)

	creationTime, err := time.Parse(time.RFC3339, request.CreationTime)
	if err != nil {
		return err
	}

	requestInJson, err := json.Marshal(request)
	if err != nil {
		return err
	}

	_, err = conn.Do("setnx", key, string(requestInJson))
	if err != nil {
		return err
	}

	_, err = conn.Do("zadd", "ids", "nx", creationTime.Unix(), key)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) GetDeliveries(flushdb bool) (*[]api.Delivery, error) {
	rawValues, err := s.GetRawValues(flushdb)
	if err != nil {
		return nil, err
	}

	deliveries := make([]api.Delivery, len(*rawValues))
	for i, v := range *rawValues {
		var delivery api.Delivery

		err = json.Unmarshal([]byte(v), &delivery)
		if err != nil {
			return nil, err
		}

		deliveries[i] = delivery
	}

	return &deliveries, nil
}

func (s *storage) GetDelivery(deliveryId api.DeliveryID) (api.Delivery, error) {
	var (
		key      string = "delivery:" + strconv.FormatUint(uint64(deliveryId), 10)
		delivery api.Delivery
	)

	value, err := s.GetRawValue(key)
	if err != nil {
		return delivery, err
	}

	err = json.Unmarshal([]byte(value), &delivery)
	if err != nil {
		return delivery, err
	}

	return delivery, nil
}

func (s *storage) GetRawValues(flushdb bool) (*[]string, error) {
	conn := s.pool.Get()
	defer conn.Close()

	byteSlices, err := redis.ByteSlices(conn.Do("zrange", "ids", 0, time.Now().Unix()))
	if err != nil {
		return nil, err
	}

	rawValues := make([]string, len(byteSlices))
	for i, v := range byteSlices {
		rawValue, err := s.GetRawValue(string(v))
		if err != nil {
			return nil, err
		}

		rawValues[i] = rawValue
	}

	if flushdb {
		_, err = conn.Do("flushdb", "async")
		if err != nil {
			return nil, err
		}
	}

	return &rawValues, nil
}

func (s *storage) GetRawValue(key string) (string, error) {
	conn := s.pool.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("get", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}
