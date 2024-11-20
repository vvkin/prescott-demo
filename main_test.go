package main

import (
	"github.com/kellydunn/golang-geo"
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	point1 := geo.NewPoint(40.7128, -74.0060)
	point2 := geo.NewPoint(34.0522, -118.2437)
	expected := 3935.75
	result := point1.GreatCircleDistance(point2)

	if result < expected-1.0 || result > expected+1.0 {
		t.Errorf("Expected distance around %v, but got %v", expected, result)
	}
}
