package main

import "C"
import (
	"Go/Cricle"
	"Go/Rectangle"
	"bufio"
	"fmt"
	"os"
	"strings"
	"Go/Input"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for{
		    Input.Input()
		
			fmt.Println("Welcome To Math Program")
			fmt.Println("Circle?\nRectangle?")
			scanner.Scan()
			secim := strings.ToLower(scanner.Text())
			if secim == strings.ToLower("q") {
				break
			}
			if secim == "rectangle" {
				rectangle.Rectangle()
			} else if secim == "circle" {
				circle.Cricle()
			} else {
				fmt.Println("Wrong Entry")
			}

	}

}
