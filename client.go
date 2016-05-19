package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type WeatherProvider interface {
	GetWeather(cityID string) (SimpleWeatherResponse, error)
}

type HttpWeatherProvider struct {
	APIKey string
}

func timed(call func() (*http.Response, error), metricsKey string) (*http.Response, error) {
	if !metricsEnabled {
		log.Print("Skip sending metrics to DD")
		return call()
	}

	log.Print("Sending metrics to DD")
	before := time.Now().Nanosecond()
	r, e := call()
	durationMs := float64((time.Now().Nanosecond() - before) / 1000 / 1000)
	if err := metricsClient.TimeInMilliseconds(metricsKey, durationMs, []string{}, 1.0); err != nil {
		log.Printf("failed to report timing: %s\n", err.Error())
	}
	return r, e
}

func (p HttpWeatherProvider) GetWeather(cityID string) (SimpleWeatherResponse, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?id=%s&APPID=%s&units=%s", cityID, p.APIKey, "metric")

	res, err := timed(func() (*http.Response, error) { return http.Get(url) }, "response.time.ms")
	if err != nil {
		return SimpleWeatherResponse{}, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return SimpleWeatherResponse{}, err
	}
	// TODO add support for handling error JSON responses
	var data WeatherResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return SimpleWeatherResponse{}, err
	}

	return convert(data), nil
}

func getWeatherProvider() (HttpWeatherProvider, error) {
	val, err := ioutil.ReadFile("/configuration.txt")
	if err != nil {
		return HttpWeatherProvider{}, err
	}
	return HttpWeatherProvider{strings.TrimSpace(string(val))}, nil
}

func convert(res WeatherResponse) SimpleWeatherResponse {
	cond := make([]string, len(res.Weather))
	for i := 0; i < len(res.Weather); i++ {
		cond[i] = res.Weather[i].Main
	}
	return SimpleWeatherResponse{res.Main, res.Wind, res.Clouds, cond}
}
