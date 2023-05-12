package weather

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func getWeather(city_name string, country_code string) (string, string){

    apiKey := "YOUR_API_KEY" // 7b891d40f8a7353ab15108fbbf741700 apply .env to this

    url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s,%s&appid=%s", city_name, country_code, apiKey)

    
    response, err := http.Get(url)
    if err != nil {
        panic(err)
    }

    
    const weatherResponse WeatherResponse
    err = json.NewDecoder(response.Body).Decode(&weatherResponse)
    if err != nil {
        panic(err)
    }

    weather := weatherResponse.Weather[0]
    temperature := weatherResponse.Main.Temp

    return weather , temperature
}

type WeatherResponse struct {
    Weather []Weather `json:"weather"`
    Main    Main      `json:"main"`
}

type Weather struct {
    Description string `json:"description"`
}

type Main struct {
    Temp float64 `json:"temp"`
}
