package main

import (
	"encoding/json"
	"fmt"
)

type WeatherJSON struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity string  `json:"humidity"`
	} `json:"main"`
}

func decoding() {
	jsonData := `{
		"name" : "Colombo",
		"main" : {
			"temp" : 20.03,
			"humidity" : "sunny"
		}
	}`

	var unpackJson WeatherJSON

	error := json.Unmarshal([]byte(jsonData), &unpackJson)

	fmt.Println(error)

	if error != nil {
		fmt.Println("Bad JSON : ",error)
		return
	}

	fmt.Println(unpackJson.Name)
	fmt.Println(unpackJson.Main)
	fmt.Println(unpackJson.Main.Humidity)
	fmt.Println(unpackJson.Main.Temp)
}