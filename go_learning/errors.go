package main

import (
	"errors"
	"fmt"
)

func division(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("Cannot Divide by zero !")
	}else {
		return num1 / num2, nil
	}
}

func checkTemperature(temp float64) (string ,error){
	if (temp < -100 || temp > 60){
		return "", fmt.Errorf("Unrealistic temperature (%.2fC) X",temp)
	}else {
		return fmt.Sprintf("Temperature is : %.2f",temp), nil
	}
}