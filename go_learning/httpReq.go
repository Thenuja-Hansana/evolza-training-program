// Request URL : https://jsonplaceholder.typicode.com/users/1

package main

import (
	"fmt"
	"io"
	"net/http"
)

func request(){

	response, error := http.Get("https://jsonplaceholder.typicode.com/users/1")

	if error != nil {
		fmt.Println("Bad Request!",error)
	}

	defer response.Body.Close()

	dataInByte, error := io.ReadAll(response.Body)

	if error != nil{
		fmt.Println("Failed to read the data!",error)
		return
	}

	fmt.Println(string(dataInByte))
}


