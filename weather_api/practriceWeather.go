package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherResponses struct {
	Location struct {
		Name    string  `json:"name"`
		Region  string  `json:"region"`
		Country string  `json:"country"`
		Lat     float64 `json:"lat"`
		Lon     float64 `json:"lon"`
	}`json:"location"`

	Current struct {
		TempC    float64 `json:"temp_c"`
		Humidity int     `json:"humidity"`
	}`json:"current"`
}

func getCurrentWeather(city string) (WeatherResponses, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=5c0b47f0503b41db89c55209260607&q=%s", city)

	response, error := http.Get(url)

	if error != nil {
		return WeatherResponses{}, error
	}

	defer response.Body.Close()

	body, error := io.ReadAll(response.Body)

	if error != nil {
		return WeatherResponses{}, error
	}

	var weathers WeatherResponses

	err := json.Unmarshal(body, &weathers)

	if err != nil {
		return WeatherResponses{}, err
	}

	return weathers ,nil

}

func run(cityName string){
	weathers, err := getCurrentWeather(cityName)
	
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 4. Updated the print statements to match your new struct!
	fmt.Println("City:", weathers.Location.Name)
	fmt.Println("Country:", weathers.Location.Country)
	fmt.Println("Temp (C):", weathers.Current.TempC)
}
