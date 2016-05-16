package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/_ah/health", getHealth)
	http.HandleFunc("/rest/forecast", getForecast)

	fmt.Println("Initializing..")
	cacheEnabled = initializeCache()
	if err := initializeMetrics(); err != nil {
		fmt.Printf("Failed to initialize metrics client: %s", err.Error())
	}

	fmt.Println("Running..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
