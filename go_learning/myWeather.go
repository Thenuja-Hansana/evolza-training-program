package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
type WeatherResponse struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

func getWeather(city string) (WeatherResponse, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=5c0b47f0503b41db89c55209260607 &q=%s", city)

	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err   // return an EMPTY struct + the error
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherResponse{}, err
	}

	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return WeatherResponse{}, err
	}

	return weather, nil   // success: return the filled struct + no error
}

func mainR() {
	weather, err := getWeather("Colombo")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("City:", weather.Name)
	fmt.Println("Temp:", weather.Main.Temp)
}