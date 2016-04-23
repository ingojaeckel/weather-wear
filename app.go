package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", GetHealth)
	http.HandleFunc("/rest/forecast", GetForecast)

	fmt.Println("Running..")

	initializeCache()

	http.ListenAndServe(":8080", nil)
}
