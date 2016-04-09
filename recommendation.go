package main

import "fmt"

// Recommendation on what to wear on a given day
type Recommendation struct {
	sunglasses                 bool
	hat                        bool
	shorts, jeans, snowpants   bool
	tshirt, sweater, jacket    bool
	flipFlops, sneakers, boots bool
	umbrella                   bool
}

func (r Recommendation) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s",
		defaultOrEmpty(r.hat, "hat "),
		defaultOrEmpty(r.shorts, "shorts "),
		defaultOrEmpty(r.jeans, "jeans "),
		defaultOrEmpty(r.snowpants, "snowpants "),

		defaultOrEmpty(r.tshirt, "tshirt "),
		defaultOrEmpty(r.sweater, "sweater "),
		defaultOrEmpty(r.jacket, "jacket "),

		defaultOrEmpty(r.flipFlops, "flipFlops "),
		defaultOrEmpty(r.sneakers, "sneakers "),
		defaultOrEmpty(r.boots, "boots "),

		defaultOrEmpty(r.umbrella, "umbrella "),
	)
}

func defaultOrEmpty(val bool, defaultVal string) string {
	if val {
		return defaultVal
	}
	return ""
}

func GetRecommendation(weather SimpleWeatherResponse) Recommendation {
	var r Recommendation

	// Temperature
	if weather.Main.Temp < 0 {
		r.snowpants = true
		r.hat = true
		r.boots = true
		r.jacket = true
		r.sweater = true
	} else if weather.Main.Temp >= 0 && weather.Main.Temp < 10 {
		r.jacket = true
		r.sweater = true
		r.jeans = true
		r.boots = true
	} else if weather.Main.Temp >= 10 && weather.Main.Temp < 20 {
		r.sneakers = true
		r.jeans = true
		r.tshirt = true
	} else { // temp > 20
		r.shorts = true
		r.tshirt = true
		r.flipFlops = true
	}

	// Wind
	if weather.Wind.Speed > 10 {
		r.jacket = true
	}

	// Rain / Brightness
	for i := 0; i < len(weather.Conditions); i++ {
		if weather.Conditions[i] == "Clear" {
			r.sunglasses = true
			break
		}
		if weather.Conditions[i] == "Rain" {
			r.umbrella = true
			break
		}
	}
	return r
}
