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
		writeJSONResponse(w, http.StatusBadRequest, RecommendationResponse{Status: 1, Error: "Missing cityId"})
		return
	}
	if len(c) != 1 {
		writeJSONResponse(w, http.StatusBadRequest, RecommendationResponse{Status: 2, Error: fmt.Sprintf("Invalid number of cityIds: %d", len(c))})
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
		writeJSONResponse(w, http.StatusInternalServerError, RecommendationResponse{Status: 3, Error: err.Error()})
		return
	}

	resp, err := s.GetWeather(cityID)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, RecommendationResponse{Status: 4, Error: err.Error()})
		return
	}

	val := GetRecommendation(resp).String()
	writeJSONResponse(w, http.StatusOK, RecommendationResponse{Recommendation: val})
	cachePut(cityID, val, forecastCacheTTLSeconds)
}

func writeJSONResponse(w http.ResponseWriter, status int, r RecommendationResponse) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, r.String())
}
