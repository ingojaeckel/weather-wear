package main

import (
	"fmt"
	"log"

	"github.com/DataDog/datadog-go/statsd"
)

const statsdHostname = "docker-dd-agent-dot-weather-wea.appspot.com"
const statsdPort = 8125

var metricsClient *statsd.Client

func initializeMetrics() bool {

	metricsClient, err := statsd.New(fmt.Sprintf("%s:%d", statsdHostname, statsdPort))

	if err != nil {
		log.Fatal(err)
	}
	// prefix every metric with the app name
	metricsClient.Namespace = "flubber."
	// send the EC2 availability zone as a tag with every metric
	metricsClient.Tags = append(metricsClient.Tags, "us-east-1a")
	metricsClient.Gauge("request.duration", 1.2, nil, 1)

	return true
}
