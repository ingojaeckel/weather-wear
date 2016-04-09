package main

import (
	"fmt"
	"testing"
)

const (
	SanFrancisco = "5391997"
	Cleveland    = "4461030"
	NewYork      = "5128581"
	Seattle      = "5809844"
)

func TestServerInteraction(t *testing.T) {
	fmt.Println(getRecommendation(t, SanFrancisco))
	fmt.Println(getRecommendation(t, Cleveland))
	fmt.Println(getRecommendation(t, NewYork))
	fmt.Println(getRecommendation(t, Seattle))
}

func TestRecommendationSummer(t *testing.T) {
	s := SimpleWeatherResponse{
		Temperature{25.0, 100.0, 20.0, 18.0, 22.0, 19.0, 20.0},
		Wind{10.0, 0.0},
		Clouds{50},
		[]string{"Sunny"},
	}

	r := GetRecommendation(s)
	if !r.tshirt || !r.shorts || !r.flipFlops {
		t.Error("Come on, it's hot outside.")
	}
}

func TestRecommendationSpring(t *testing.T) {
	s := SimpleWeatherResponse{
		Temperature{15.0, 100.0, 20.0, 18.0, 22.0, 19.0, 20.0},
		Wind{10.0, 0.0},
		Clouds{50},
		[]string{"Sunny"},
	}

	r := GetRecommendation(s)
	if !r.tshirt || !r.jeans || !r.sneakers {
		t.Error("Come on, it's warm enough.")
	}
	if str := fmt.Sprintf("%v", r); str != "jeans tshirt sneakers " {
		t.Errorf("Got different string instead: \"%s\"\n", str)
	}
}

func TestRecommendationWinter(t *testing.T) {
	s := SimpleWeatherResponse{
		Temperature{5.0, 100.0, 20.0, 18.0, 22.0, 19.0, 20.0},
		Wind{10.0, 0.0},
		Clouds{50},
		[]string{"Sunny"},
	}

	r := GetRecommendation(s)
	//jeans tshirt sneakers
	if !r.sweater || !r.jeans || !r.boots {
		t.Error("Come on, it's warm enough.")
	}
	if str := fmt.Sprintf("%v", r); str != "jeans sweater jacket boots " {
		t.Errorf("Got different string instead: \"%s\"\n", str)
	}
}

func getRecommendation(t *testing.T, cityId string) Recommendation {
	s, err := getWeatherProvider()
	if err != nil {
		t.Skip("Skipping test since WeatherProvider could not be created. This might be due to missing configuration.")
	}
	res, err := s.GetWeather(cityId)
	if err != nil {
		t.Errorf("Failed to download weather: \"%f\"\n", err.Error())
	}
	if res.Main.Temp == 0.0 || res.Main.Temp < -40 && res.Main.Temp > 50 {
		t.Errorf("Temperature value looks incorrect.")
	}
	return GetRecommendation(res)
}
