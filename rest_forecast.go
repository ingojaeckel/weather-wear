package main

import (
	"fmt"
	"io"
	"net/http"
)

func getForecast(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	c, ok := q["cityId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		result := RecommendationResponse{Status: 1, Error: "Missing cityId"}
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, result.String())
		return
	}
	if len(c) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		result := RecommendationResponse{Status: 2, Error: fmt.Sprintf("Invalid number of cityIds: %d", len(c))}
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, result.String())
		return
	}

	cityID := c[0]
	cached, err := cacheGet(cityID)
	if cacheEnabled && err == nil && len(cached) > 0 {
		io.WriteString(w, cached)
		return
	}

	s, err := getWeatherProvider()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		result := RecommendationResponse{Status: 3, Error: err.Error()}
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, result.String())
		return
	}

	resp, err := s.GetWeather(cityID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		result := RecommendationResponse{Status: 4, Error: err.Error()}
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, result.String())
		return
	}

	result := RecommendationResponse{Recommendation: GetRecommendation(resp).String()}
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, result.String())
	cachePut(cityID, result.String(), forecastCacheTTLSeconds)
}
