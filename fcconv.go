package main

import "fmt"

func main() {
	var temp float64
	var unit string

	fmt.Println("'F' = Fahrenheit")
	fmt.Println("'C' = Celsius")

	fmt.Println("What Unit Would You Like To Convert To: ")
	fmt.Scan(&unit)

	if unit == "F" || unit == "f" {
		fmt.Print("Enter Tempeature In Celsius: \n")
		fmt.Scan(&temp)

		temp = (1.8 * temp) + 32.0
		fmt.Println("Tempeature Is ", temp, "F")
	} else if unit == "C" || unit == "c" {
		fmt.Print("Enter Tempeature In Fahrenheit: \n")
		fmt.Scan(&temp)

		temp = (temp - 32) / 1.8
		fmt.Println("Tempeature Is ", temp, "C")
	} else {
		fmt.Println("Invalid Input")
	}

}
