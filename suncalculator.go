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
func CalcLeftQuad(trueLong float64) float64 {
	var ninety float64 = 90.0
	return math.Floor(trueLong/ninety) * ninety
}
func CalcRightQuad(RA float64) float64 {
	var ninety float64 = 90.0
	return math.Floor(RA/ninety) * ninety
}
func RA_ToQuad(RA float64, trueLong float64) float64 {
	return RA + (CalcLeftQuad(trueLong) - CalcRightQuad(RA))
}
func RightAscenion(trueLong float64) float64 {
	var sunTan float64 = math.Tan(trueLong * math.Pi / 180)
	var innerTan = .91764 * sunTan

	var RA float64 = math.Atan(innerTan) * (180 / math.Pi)
	RA = RA_ToQuad(RA, trueLong) / 15.0
	if RA > 360 {
		RA -= 360
	} else if RA < 0 {
		RA += 360
	}
	return RA
}
func sinDec(trueLong float64) float64 {
	return .39782 * math.Sin(trueLong*math.Pi/180)
}
func cosDec(trueLong float64) float64 {
	SinDec := sinDec(trueLong)
	arcSin := math.Asin(SinDec * 180 / math.Pi)
	return math.Cos(arcSin * math.Pi / 180)
}
