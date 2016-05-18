package main

import (
	"log"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/_ah/health", getHealth)
	http.HandleFunc("/rest/forecast", getForecast)

	log.Printf("GAE IsDevAppServer: %v\n", appengine.IsDevAppServer())
	log.Println("Initializing..")
	cacheEnabled = initializeCache()
	if err := initializeMetrics(); err != nil {
		log.Printf("Failed to initialize metrics client: %s", err.Error())
	} else {
		log.Println("Initialized metrics")
	}
	if err := metricsClient.Count("some.counter", 1, []string{}, 1.0); err != nil {
		log.Printf("Failed to report counter: %s\n", err.Error())
	}
	if err := metricsClient.TimeInMilliseconds("some.timer", 1, []string{}, 1.0); err != nil {
		log.Printf("Failed to report timer: %s\n", err.Error())
	}
	log.Println("Running..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
