package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/_ah/health", getHealth)
	http.HandleFunc("/rest/forecast", getForecast)

	fmt.Println("Running..")

	cacheEnabled = initializeCache()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
