package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
func main() {
	fmt.Println("Test")
	http.HandleFunc("/test", testJson)
	http.ListenAndServe(":8090", nil)
}
