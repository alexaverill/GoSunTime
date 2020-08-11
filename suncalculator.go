package main

import (
	"math"
)

func calculateDayOfYear(month int, day int, year int) int {
	var one = math.Floor((float64(275) * float64(month) / float64(9.0)))
	var two = math.Floor((float64(month) + float64(9)) / float64(12.0))
	var three = (1 + math.Floor((float64(year)-float64(4)*math.Floor(float64(year)/float64(4))+float64(2))/float64(3)))
	return int(one) - int((two * three)) + day - 30
}
func LongitudeHourToTime(day float64, lngHour float64, sunset bool) float64 {
	if sunset {
		return day + ((float64(18) - lngHour) / float64(24))
	}
	return day + ((6 - lngHour) / 24)
}
func meanAnomaly(time float64) float64 {
	return (.9856 * time) - 3.289
}
func SunLongitude(meanAnomaly float64) float64 {
	var lng = meanAnomaly + (1.916 * math.Sin(meanAnomaly*math.Pi/180)) + (.02*math.Sin(2*meanAnomaly*math.Pi/180) + 282.634)
	if lng > 360 {
		return lng - 360
	}
	if lng < 0 {
		return lng + 360

	}
	return lng
}
