package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type StdError struct {
	ErrorMessage string
}

func test(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Test\n")
}
func testJson(w http.ResponseWriter, req *http.Request) {

	var sunEvent SunEventInfo = CalculateSunTime(6, 25, 1990, -.01454, 40.9, -74.3, false)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sunEvent)

	//fmt.Fprintf(w,"Test\n");
}
func calculate(req *http.Request, sunset bool) SunEventInfo {
	vals := req.URL.Query()
	//TODO: add error handling and validation
	month, _ := strconv.Atoi(vals.Get("month"))
	day, _ := strconv.Atoi(vals.Get("day"))
	year, _ := strconv.Atoi(vals.Get("year"))
	lat, _ := strconv.ParseFloat(vals.Get("lat"), 64)
	lng, _ := strconv.ParseFloat(vals.Get("lng"), 64)
	return CalculateSunTime(month, day, year, -.01454, lat, lng, sunset)
}
func sunrise(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")
	var sunEvent SunEventInfo = calculate(req, false)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sunEvent)
}
func sunset(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")
	var sunEvent SunEventInfo = calculate(req, true)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sunEvent)
}

func main() {
	fmt.Println("Test")
	http.HandleFunc("/test", testJson)
	http.HandleFunc("/sunset", sunset)
	http.HandleFunc("/sunrise", sunrise)
	http.ListenAndServe(":8090", nil)
}
