package main

import (
	"fmt"
	"io"
	"net/http"
)

// TODO return JSON, return proper HTTP status codes
func getForecast(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	c, ok := q["cityId"]
	if !ok {
		io.WriteString(w, fmt.Sprintf("Missing multiple cityId"))
		return
	}
	if len(c) != 1 {
		io.WriteString(w, fmt.Sprintf("Invalid number of cityIds: %d", len(c)))
		return
	}

	cityID := c[0]
	cached, err := cacheGet(cityID)
	if err == nil && len(cached) > 0 {
		io.WriteString(w, cached)
		return
	}

	s, err := getWeatherProvider()
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Failed: %s", err.Error()))
		return
	}

	resp, err := s.GetWeather(cityID)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Failed: %s", err.Error()))
		return
	}

	result := GetRecommendation(resp).String()
	io.WriteString(w, result)
	cachePut(cityID, result)
}
