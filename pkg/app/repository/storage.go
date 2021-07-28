package repository

import (
	"github.com/go-redis/redis"
)

type Storage interface {
	CreateDelivery()
	GetDeliveries()
	GetDelivery()
}

type storage struct {
	db *redis.Client
}

func NewStorage() Storage {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &storage{db: redis}
}

func (s *storage) CreateDelivery() {}
func (s *storage) GetDeliveries()  {}
func (s *storage) GetDelivery()    {}
