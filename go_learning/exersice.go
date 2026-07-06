package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Target URL: https://jsonplaceholder.typicode.com/todos/1

/*
	{
		"userId": 1,
		"id": 1,
		"title": "delectus aut autem",
		"completed": false
	}
*/

type Todo struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func students() {

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil{
		fmt.Println("Bad Request!",err)
		return
	}

	defer response.Body.Close()

	dataInByte, error := io.ReadAll(response.Body)
	if error != nil{
		fmt.Println("Issue occur when reading the data!",error)
		return
	}

	//fmt.Println(string(dataInByte))

	var todoData Todo

	error = json.Unmarshal(dataInByte, &todoData)
	if error != nil{
		fmt.Println("Issue occur while parsing the JSON!",error)
		return
	}

	fmt.Println("Title : ", todoData.Title)
	fmt.Println("Completed : ",todoData.Completed)

}

