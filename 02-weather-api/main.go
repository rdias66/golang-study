package main

import (
	"./src/weather.service"
	"./src/user.service"
	"./src/printer.service"
)

func main () {
	city_name , country_code := inputInfo()

	weather_response , temperature_response := getWeather(city_name, country_code) 
	
	printer(weather_response, temperature_response)
}
