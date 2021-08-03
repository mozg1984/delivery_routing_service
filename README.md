# Test web service (delivery routing service)

It provides 4 endpoints:
- POST     /deliveries/add_delivery     (it creates a new delivery item)
- GET      /deliveries/list             (it returns the list of delivery items with flushing them on the persist layer)
- GET      /deliveries/:id              (it returns the delivery item by given id)
- GET      /deliveries/route_distance   (it calculates the route distance by using all delivery items)

By default, the server listens for requests on the address - `0.0.0.0:8080`

## Installation

It requires [Golang](https://golang.org/) v1.16.6+ and [Redis](https://redis.io/) v.6.2.1+ to run.

Install the dependencies, run tests and start the server.

```sh
cd delivery_routing_service
go build ./...
go test ./...
go run cmd/server.go
```
