package main

import "fmt"

func variables() {
	varDecExplicit()
	varDecShort()

	var count int
	fmt.Println(count) // 0

	const apiKey = "8s9q3dss"
}

func varDecExplicit() {
	var city string = "Ambalangoda"
	var temp float64 = 30.5
	var isRainning bool = true
	city = "Galle"

	fmt.Printf("My city is : %s\n", city)
	fmt.Printf("The tempreature is : %f\n", temp)
	fmt.Printf("Raining : %t\n", isRainning)

}

func varDecShort() {
	fname := "Thenuja"
	lname := "Hansana"
	fmt.Printf("Hi, My name is %s %s\n", fname, lname)
}
