package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (p HttpWeatherProvider) GetWeather(cityID string) (SimpleWeatherResponse, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?id=%s&APPID=%s&units=%s", cityID, p.APIKey, "metric")

	before := time.Now().Nanosecond()
	res, err := http.Get(url)
	if err != nil {
		return SimpleWeatherResponse{}, err
	}
	durationMs := float64((time.Now().Nanosecond() - before) / 1000 / 1000)
	metricsClient.TimeInMilliseconds("response.time.ms", durationMs, []string{}, 1.0)
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
	val, err := ioutil.ReadFile("configuration.txt")
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
