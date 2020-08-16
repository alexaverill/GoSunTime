package main

import (
	"encoding/json"
	"fmt"
	"log"
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
func calculate(req *http.Request, sunset bool) (SunEventInfo, bool) {
	vals := req.URL.Query()
	//TODO: add error handling and validation
	month, monthError := strconv.Atoi(vals.Get("month"))
	day, dayError := strconv.Atoi(vals.Get("day"))
	year, yearError := strconv.Atoi(vals.Get("year"))
	lat, latERror := strconv.ParseFloat(vals.Get("lat"), 64)
	lng, lngError := strconv.ParseFloat(vals.Get("lng"), 64)
	if month < 0 || month > 12 || monthError != nil {
		return SunEventInfo{}, true

	}
	if day < 0 || day > 31 || dayError != nil {
		return SunEventInfo{}, true
	}
	if yearError != nil {
		return SunEventInfo{}, true
	}
	if latERror != nil || lngError != nil {
		return SunEventInfo{}, true
	}

	return CalculateSunTime(month, day, year, -.01454, lat, lng, sunset), false
}
func sunrise(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")
	sunEvent, err := calculate(req, false)
	if err {
		log.Println("Error Parsing Sunrise request")
		log.Println(req.URL.String())
		errorMessage := StdError{ErrorMessage: "Arguements invalid! Expecting: month:int [0..12], day:int [0..31], year:int, lat:float, lng:float"}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sunEvent)
}
func sunset(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")
	sunEvent, err := calculate(req, true)
	if err {
		log.Println("Error Parsing Sunrise request")
		log.Println(req.URL.String())
		errorMessage := StdError{ErrorMessage: "Arguements invalid! Expecting: month:int [0..12], day:int [0..31], year:int, lat:float, lng:float"}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sunEvent)
}

func main() {
	http.HandleFunc("/test", testJson)
	http.HandleFunc("/sunset", sunset)
	http.HandleFunc("/sunrise", sunrise)
	http.ListenAndServe(":5000", nil)
}
