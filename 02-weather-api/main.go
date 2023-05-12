package main

import (
	"github.com/rdias66/golang-study/02-weather-api/src/printer"
	"github.com/rdias66/golang-study/02-weather-api/src/user"
	"github.com/rdias66/golang-study/02-weather-api/src/weather"
)

func main() {
	city_name, country_code := user.InputInfo()

	weather_response, temperature_response := weather.GetWeather(city_name, country_code)

	printer.Printer(weather_response, temperature_response)
}
