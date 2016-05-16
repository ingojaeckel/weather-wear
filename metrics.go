package main

import (
	"fmt"

	"github.com/DataDog/datadog-go/statsd"
)

const statsdHostname = "docker-dd-agent-dot-weather-wea.appspot.com"
const statsdPort = 8125

var metricsClient *statsd.Client

func initializeMetrics() error {
	metricsClient, err := statsd.New(fmt.Sprintf("%s:%d", statsdHostname, statsdPort))

	if err != nil {
		metricsEnabled = false
		return err
	}
	// prefix every metric with the app name
	metricsClient.Namespace = "dev."
	metricsClient.Tags = append(metricsClient.Tags, fmt.Sprintf("appid:weather-wea"))
	metricsClient.SimpleEvent("initialized", "datadog has been initialized")

	metricsEnabled = true
	return nil
}
