package main

import "fmt"

type Weather struct {
	City        string
	Temperature float64
	Condition   string
}

func structs(){
	w := Weather {
		City : "Galle",
		Temperature : 30.5,
		Condition : "Sunny",
	}
	
	fmt.Println(w.printStructs())
	
	
}

func (data  Weather) printStructs() string{
	return fmt.Sprintf("The weather in %s is %.2f and seems pretty %s",data.City, data.Temperature, data.Condition)
}
