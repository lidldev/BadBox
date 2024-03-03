package main

import "fmt"

func main() {
	var rows int
	var columns int
	var symbol string

	fmt.Println("How Many Rows: ")
	fmt.Scan(&rows)

	fmt.Println("How Many Columns: ")
	fmt.Scan(&columns)

	fmt.Println("Which Symbol: ")
	fmt.Scan(&symbol)

	for i := 1; i <= rows; i++ {
		for j := 1; j <= columns; j++ {
			fmt.Print(symbol)
		}
		fmt.Print("\n")
	}
}
