package main

import "fmt"

func main() {
	s, err := getWeatherProvider()
	if err != nil {
		fmt.Printf("Failed: %s", err.Error())
		return
	}

	// San Francisco
	resp, err := s.GetWeather("5391997")
	if err != nil {
		fmt.Printf("Failed: %s", err.Error())
		return
	}

	fmt.Printf("Temp (C):   %f\n", resp.Main.Temp)
	fmt.Printf("Conditions: %s\n", resp.Conditions)

	fmt.Printf("Recommendation: %v\n", GetRecommendation(resp))
}
