package user

import (
	"fmt"
)

func InputInfo() (string, string) {
	var city_name string
	var country_code string

	fmt.Println("Enter the city name: ")
	fmt.Scan(&city_name)

	fmt.Println("Enter the country code: ")
	fmt.Scan(&country_code)

	return city_name, country_code
}
