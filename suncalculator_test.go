package main

import (
	"fmt"
	"math"
	"testing"
)

func isEqualEnough(a float64, b float64) bool {
	var absA = math.Abs(a)
	var absB = math.Abs(b)
	var diff = math.Abs(a - b)
	var min float64 = 1e-9
	var e float64 = .001

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

func TestSunLong(t *testing.T) {
	var mean = 170.626
	var expected = 93.56
	var actual = SunLongitude(mean)
	if !isEqualEnough(actual, expected) {
		t.Errorf("Sun Longitude is wrong: %f", actual)
	}
}
func TestCalcLeftQuad(t *testing.T) {
	var expected float64 = 90
	var sunLong float64 = 93.56
	var actual = CalcLeftQuad(sunLong)
	fmt.Println(actual)
	if !isEqualEnough(actual, expected) {
		t.Errorf("Expected: %f Actual: %f", actual, expected)
	}
}
func TestCalcRightQuad(t *testing.T) {
	var expected float64 = -90
	RA := -86.11412
	actual := CalcRightQuad(RA)
	if !isEqualEnough(actual, expected) {
		t.Errorf("Expected: %f Actual: %f", actual, expected)
	}
}
func TestRA_ToQuad(t *testing.T) {
	var expected = 93.886
	RA := -86.11412
	sunLong := 93.56
	actual := RA_ToQuad(RA, sunLong)
	if !isEqualEnough(actual, expected) {
		t.Errorf("Expected: %f Actual: %f", actual, expected)
	}
}
func TestRightAscenion(t *testing.T) {
	sunLong := 93.56
	expected := 6.259
	actual := RightAscenion(sunLong)
	if !isEqualEnough(actual, expected) {
		t.Errorf("Right Accension is incorrect actual: %f Expected: %f", actual, expected)
	}

}
func TestSinDec(t *testing.T) {
	var sunLong = 93.56
	expected := .39705
	actual := sinDec(sunLong)
	if !isEqualEnough(actual, expected) {
		t.Errorf("Sin dec is actually : %f", actual)
	}
}
func TestCosDec(t *testing.T) {
	var sunLong = 93.56
	expected := .91780
	actual := cosDec(sunLong)
	if !isEqualEnough(actual, expected) {
		t.Errorf("Sin dec is actually : %f", actual)
	}
}
func TestCalcSunHour(t *testing.T) {
	localHour := -.39570
	expected := 16.446
	actual := CalculateSunHour(localHour, false)
	fmt.Printf("%f", actual)
	if !isEqualEnough(actual, expected) {
		t.Errorf("Actual: %f Expected: %f", actual, expected)
	}
}
func TestLocalHourAngle(t *testing.T) {
	lat := 40.9
	lng := -74.3
	zenith := -0.01454
	trueLon := 93.56
	actual := LocalHourAngle(zenith, lat, lng, trueLon)
	expected := -.39314
	if !isEqualEnough(actual, expected) {
		t.Errorf("Actual: %f Expected: %f", actual, expected)
	}

}
func TestCalcSunTime(t *testing.T) {
	day := 25
	month := 6
	year := 1990
	lat := 40.9
	lng := -74.3
	zenith := -0.01454
	actual := CalculateSunTime(month, day, year, zenith, lat, lng, false)
	expected := SunEventInfo{hour: 9, minute: 26}
	fmt.Println(actual.hour)
	fmt.Println(actual.minute)
	if actual.hour != expected.hour && actual.minute != expected.minute {
		t.Errorf("Times do not match!")
	}
}
func Test_cosDec(t *testing.T) {
	type args struct {
		trueLong float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cosDec(tt.args.trueLong); got != tt.want {
				t.Errorf("cosDec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_floatTimeToSplit(t *testing.T) {
	type args struct {
		time float64
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{
			name:  "test1",
			args:  args{time: 9.5},
			want:  9,
			want1: 30,
		},
		{
			name:  "test2",
			args:  args{time: 9.44},
			want:  9,
			want1: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := floatTimeToSplit(tt.args.time)
			if got != tt.want {
				t.Errorf("floatTimeToSplit() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("floatTimeToSplit() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
