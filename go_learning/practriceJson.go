package main

import (
	"encoding/json"
	"fmt"
)

type SimpleWeather struct {
	City      string  `json:"city"`
	Temp      float64 `json:"temp"`
	Condition string  `json:"condition"`
}

func decodingWeather() {

	jsonData := `{
		"city" : "Kandy",
		"temp" : 25.3,
		"condition" : "Cloudy"
	}`

	var weatherData SimpleWeather

	err := json.Unmarshal([]byte(jsonData), &weatherData)

	if err != nil {
		fmt.Println("JSON Error :", err)
		return
	}

	fmt.Println(weatherData.City)
	fmt.Println(weatherData.Condition)
	fmt.Println(weatherData.Temp)
}