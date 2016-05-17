package main

import (
	"fmt"
	"log"

	"github.com/DataDog/datadog-go/statsd"
)

// External IP address of docker-dd-agent instance determined by running the following
// curl -H "Metadata-Flavor: Google" http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/access-configs/0/external-ip
const statsdHostname = "146.148.87.97"
const statsdPort = 8125

var metricsClient *statsd.Client

func initializeMetrics() error {
	metricsClient, err := statsd.New(fmt.Sprintf("%s:%d", statsdHostname, statsdPort))

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
