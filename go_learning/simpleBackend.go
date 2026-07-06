package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	1. Create the blue print
	2. Create the handler
*/


type UserProfile struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Age int `json:"age"`
}

// Handler
// w - writer {direct pipline back to the user}
// r - request {contains everthing that user sent to the server}
func profileHandler(w http.ResponseWriter, r *http.Request){
	
	// Telling the browser we are sending json data
	w.Header().Set("Content-Type","application/json")

	// Create Go data 
	myUser := UserProfile{
		Name : "Thenuja",
		Role : "Software Engineer",
		Age : 19,
	}

	// Go -> JSON
	error := json.NewEncoder(w).Encode(myUser)

	if error != nil {
		fmt.Println("Error occur while encoding the data to json", error)
		return
	}

}

func server(){
	// Connect the URL to the Handler 
	http.HandleFunc("/profile",profileHandler)

	// Start the server 
	fmt.Println("Server is Running\n Open : http://localhost:8080/profile")

	error := http.ListenAndServe(":8080",nil)

	if error != nil{
		fmt.Println("Server Crashed!",error)
	}
}

