package main

import (
	"math"
	"testing"
)

func isEqualEnough(a float64, b float64) bool {
	tolerance := .0001
	diff := math.Abs(a - b)
	if diff < tolerance {
		return true
	}
	return false
}
func TestCalculateDayOfYear(t *testing.T) {
	var day = calculateDayOfYear(6, 25, 1990)
	if day != 176 {
		t.Errorf("Wrong Day")
	}
	var newDay = calculateDayOfYear(8, 10, 2020)
	if newDay != 223 {
		t.Errorf("Day is: %d", newDay)
	}
}
func TestLongitudeHour(t *testing.T) {
	var lngHour = (-74.3) / 15
	var actual = LongitudeHourToTime(176, lngHour, false)
	if isEqualEnough(actual, 176.456) {
		t.Errorf("time is actually: %f", actual)
	}
}
