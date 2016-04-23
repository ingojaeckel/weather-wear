package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/_ah/health", getHealth)
	http.HandleFunc("/rest/forecast", getForecast)

	fmt.Println("Running..")

	initializeCache()

	http.ListenAndServe(":8080", nil)
}
