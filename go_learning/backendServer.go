package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CarCatalog struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year int `json:"year"`
}


func carHandler(w http.ResponseWriter, r *http.Request){

	// tell the browser we are sending json
	w.Header().Set("Content-Type","application/json")

	// data
	myCar := CarCatalog{
		Brand: "BMW",
		Model: "M4",
		Year: 2024,
	}

	error := json.NewEncoder(w).Encode(myCar)

	if error != nil {
		fmt.Println("Error while encoding to json!",error)
		return
	}

}

func carServer(){
	// Connect the URL to the handler
	http.HandleFunc("/profile",carHandler)

	// Start the server
	fmt.Println("Server is Running! \n Open in : http://localhost:8080/profile")

	// This will pause the terminal forever, listining for request
	error := http.ListenAndServe(":8080",nil)

	if error != nil{
		fmt.Println("Server Crashed!",error)
	}
}