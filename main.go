package main

import (
	"fmt"
	"github.com/kellydunn/golang-geo"
)

func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	point1 := geo.NewPoint(lat1, lon1)
	point2 := geo.NewPoint(lat2, lon2)

	distance := point1.GreatCircleDistance(point2)
	return distance
}

func main() {
	lat1, lon1 := 40.7128, -74.0060
	lat2, lon2 := 34.0522, -118.2437
	distance := calculateDistance(lat1, lon1, lat2, lon2)
	fmt.Printf("The distance between the points is: %.2f kilometers\n", distance)
}
