package main

import (
	"./src/weather.service"
)

func main () {
	city_name := "User input for city"
	country_code := "User input for country"
	weather_response , temperature_response := getWeather(city_name, country_code) 
}
