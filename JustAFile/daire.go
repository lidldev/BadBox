package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const PI = 3.14

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Area/Circumference finder (Q to quit)")

	for true {
		fmt.Println("Area or circumference:")
		scanner.Scan()
		secim := strings.ToLower(scanner.Text())

		if secim == "q" {
			break
		}

		fmt.Println("Enter the radius for circle:")
		scanner.Scan()
		yaricap, _ := strconv.ParseFloat(scanner.Text(), 64)

		if secim == "area" {
			fmt.Println(PI * (math.Pow(yaricap, 2.0)))
		} else if secim == "circumference" {
			fmt.Println(PI * 2 * yaricap)
		} else {
			fmt.Println("Invalid input!")
		}
	}
}
