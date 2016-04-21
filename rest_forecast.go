package main

import (
  "fmt"
  "io"
  "net/http"
)

// TODO return JSON, return proper HTTP status codes
func GetForecast(w http.ResponseWriter, r *http.Request) {
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

  s, err := getWeatherProvider()
  if err != nil {
    io.WriteString(w, fmt.Sprintf("Failed: %s", err.Error()))
    return
  }

  resp, err := s.GetWeather(c[0])
  if err != nil {
    io.WriteString(w, fmt.Sprintf("Failed: %s", err.Error()))
    return
  }

  io.WriteString(w, fmt.Sprint(GetRecommendation(resp)))
}
