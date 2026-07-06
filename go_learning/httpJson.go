package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
{
  "id": 1,
  "name": "Leanne Graham",
  "username": "Bret",
  "email": "Sincere@april.biz",
  "address": {
    "street": "Kulas Light",
    "suite": "Apt. 556",
    "city": "Gwenborough",
    "zipcode": "92998-3874",
    "geo": {
      "lat": "-37.3159",
      "lng": "81.1496"
    }
  },
  "phone": "1-770-736-8031 x56442",
  "website": "hildegard.org",
  "company": {
    "name": "Romaguera-Crona",
    "catchPhrase": "Multi-layered client-server neural-net",
    "bs": "harness real-time e-markets"
  }
}
*/

type PlaceHolder struct {

	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Address struct {
		Street string `json:"street"`
		City string `json:"city"`
	}
	Phone string `json:"phone"`
	Company struct {
		CompanyName string `json:"name"`
	}
}


func httpJson() {

	response, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil{
		fmt.Println("Bad Request!",err)
		return
	}

	defer response.Body.Close()

	dataInByte, error := io.ReadAll(response.Body)
	if error != nil{
		fmt.Println("Failed to read data!",error)
		return
	}

	var data PlaceHolder

	error = json.Unmarshal(dataInByte, &data)
	if error != nil{
		fmt.Println("Failed to decode the data!",error)
		return
	}
	fmt.Println("Success!")
	fmt.Println("ID : ", data.Id)
	fmt.Println("Name : ", data.Name)
	fmt.Println("Email : ", data.Email)
	fmt.Printf("Address : %s, %s\n",data.Address.Street,data.Address.City)
	fmt.Println("Phone : ",data.Phone)
	fmt.Println("Company Name : ",data.Company.CompanyName)
}

