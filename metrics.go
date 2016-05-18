package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/DataDog/datadog-go/statsd"
)

var metricsClient *statsd.Client

func initializeMetrics() error {
	statsdAddress := getStatsdAddress()
	log.Printf("Using statsd address: %s\n", statsdAddress)
	metricsClient, err := statsd.New(statsdAddress)

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

func getStatsdAddress() string {
	res, err := http.Get("https://docker-dd-agent-dot-weather-wea.appspot.com/")
	if err != nil {
		return err.Error()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}
