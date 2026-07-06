package main

import "fmt"

func functions() {
	message := greet("Thenuja")
	fmt.Println(message)

	total := getSum(12, 22)
	fmt.Printf("Total : %d\n", total)

	divided := divide(10.2, 2.2)
	if divided == -1 {
		fmt.Println("Zero divition Error..")
	} else {
		fmt.Printf("divided : %.2f\n", divided)
	}

	weather := describeweather("Galle", 12.33)
	fmt.Println(weather)
}

func greet(name string) string {
	return "Hello, My name is " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func divide(num1 float64, num2 float64) float64 {
	if num1 == 0 {
		return -1.0
	} else {
		return num1 / num2
	}
}

func describeweather(city string, temp float64) string {
	return fmt.Sprintf("The weather in %s is %.2fC", city, temp)
}
