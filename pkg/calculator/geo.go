package geo

import (
	"math"
)

const R float64 = 6371e3 // Radius of the earth, m

func convertToRadians(deg float64) float64 {
	return deg * math.Pi / 180
}

func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	dLat := convertToRadians(lat2 - lat1)
	dLon := convertToRadians(lon2 - lon1)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(convertToRadians(lat1))*math.Cos(convertToRadians(lat2))*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return c * R
}
