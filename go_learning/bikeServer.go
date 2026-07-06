package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bike struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func bikeHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")

	myBike := Bike{
		Brand : "Yamaha",
		Model: "R1",
		Year: 2006,
	}

	error := json.NewEncoder(w).Encode(myBike)

	if error != nil{
		fmt.Println("Error occured while encoding to JSON!",error)
		return
	}
}

func bikeServer(){
	http.HandleFunc("/profile", bikeHandler)

	fmt.Println("Server is running!\n Open in : http://localhost:8080/profile")

	error := http.ListenAndServe(":8080",nil)

	if error != nil{
		fmt.Println("Server Crashed!",error)
	}
}