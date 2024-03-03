package main

import "fmt"

func main() {
	var rows int
	var columns int
	var char string

	fmt.Println("How Many Rows?")
	fmt.Scan(&rows)

	fmt.Println("How Many Columns")
	fmt.Scan(&columns)

	fmt.Println("Enter A Symbol To Use")
	fmt.Scan(&char)

	for i := 0; i <= rows; i++ {
		for j := 0; j <= columns; j++ {
			fmt.Print(char)
		}
		fmt.Print("\n")
	}

}
