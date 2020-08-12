package main

import (
	"math"
	"testing"
)

func isEqualEnough(a float64, b float64) bool {
	var absA = math.Abs(a)
	var absB = math.Abs(b)
	var diff = math.Abs(a - b)
	var min float64 = 1e-9
	var e float64 = 1e-5
	if a == b {
		return true
	}
	if a == 0 || b == 0 || (absA+absB < min) {
		return diff < (e * min)
	}
	return diff/math.Min((absA+absB), 1e10) < e
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
	output := isEqualEnough(actual, 176.456)
	t.Log(output)
	if !isEqualEnough(actual, 176.45638) {
		t.Errorf("time is actually: %f", actual)
	}
}

func testSunLong(t *testing.T) {
	var mean = 176.456
	var expected = 93.56
	var actual = SunLongitude(mean)
	if !isEqualEnough(actual, expected) {
		t.Errorf("Sun Longitude is wrong: %f", actual)
	}
}
