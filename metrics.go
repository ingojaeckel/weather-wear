package main

import (
	"fmt"
	"log"

	"github.com/DataDog/datadog-go/statsd"
)

var metricsClient *statsd.Client

func initializeMetrics() error {
	statsdAddress := "127.0.0.1:8125"
	log.Printf("Using statsd address: %s\n", statsdAddress)
	m, err := statsd.New(statsdAddress)
	metricsClient = m

	if err != nil {
		log.Printf("Disabled metrics due to error: %s\n", err.Error())
		metricsEnabled = false
		return err
	}
	log.Print("Enabled metrics")
	// prefix every metric with the app name
	metricsClient.Namespace = "dev."
	metricsClient.Tags = append(metricsClient.Tags, fmt.Sprintf("appid:weather-wea"))
	metricsClient.SimpleEvent("initialized", "datadog has been initialized")

	metricsEnabled = true
	return nil
}
