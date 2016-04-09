package main

type Coordinates struct {
	Lon, Lat float64
}

type Weather struct {
	Id                      int
	Main, Description, Icon string
}

type Temperature struct {
	Temp, Pressure, Humidity float64
	TempMin                  float64 `json:"temp_min"`
	TempMax                  float64 `json:"temp_max"`
	SeaLevel                 float64 `json:"sea_level"`
	GrndLevel                float64 `json:"grnd_level"`
}

type Wind struct {
	Speed, Deg float64
}

type Clouds struct {
	All int
}

type Sys struct {
	Message         float64
	Country         string
	Sunrise, Sunset int64
}

type WeatherResponse struct {
	Coord   Coordinates
	Weather []Weather
	Main    Temperature
	Wind    Wind
	Clouds  Clouds
	Sys     Sys

	Id, Dt, Cod int64
	Base, Name  string
}

type SimpleWeatherResponse struct {
	Main       Temperature
	Wind       Wind
	Clouds     Clouds
	Conditions []string
}
