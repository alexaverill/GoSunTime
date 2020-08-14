package main

import (
	"math"
)

type SunEventInfo struct {
	Hour, Minute int
	TimeZone     string
}

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
func MeanAnomaly(time float64) float64 {
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
	arcSin := math.Asin(SinDec) * 180 / math.Pi
	return math.Cos(arcSin * math.Pi / 180)
}
func LocalHourAngle(zenith, latitude, longitude, trueLong float64) float64 {
	sinLat := sinDec(trueLong) * math.Sin(latitude*math.Pi/180)
	numerator := zenith - sinLat
	denominator := cosDec(longitude) * math.Cos(latitude*math.Pi/180)
	angle := numerator / denominator
	return angle
}
func CalculateSunHour(localHour float64, sunset bool) float64 {
	var hour float64
	if sunset {
		hour = math.Acos(localHour) * 180 / math.Pi

	} else {
		hour = 360 - math.Acos(localHour)*180/math.Pi
	}
	hour /= 15
	return hour
}
func floatTimeToSplit(time float64) (int, int) {
	var hours int = int(time)
	var minFrac float64 = time - float64(hours)
	var minutes int = int(60 * minFrac)
	return hours, minutes
}
func CalculateSunTime(month, day, year int, zenith, lat, lng float64, sunset bool) SunEventInfo {
	dayOfYear := calculateDayOfYear(month, day, year)
	lngHour := lng / 15
	time := LongitudeHourToTime(float64(dayOfYear), lngHour, sunset)
	anomaly := MeanAnomaly(time)
	trueLongitude := SunLongitude(anomaly)
	RA := RightAscenion(trueLongitude)
	localHour := LocalHourAngle(zenith, lat, lng, trueLongitude)
	var hour float64 = CalculateSunHour(localHour, sunset)

	localEventTime := (hour + RA) - (.06571 * time) - 6.62
	UTC := localEventTime - lngHour
	if UTC < 0 {
		UTC += 24
	} else if UTC > 24 {
		UTC -= 24
	}
	hours, minutes := floatTimeToSplit(UTC)
	return SunEventInfo{Hour: hours, Minute: minutes, TimeZone: "UTC"}

}
