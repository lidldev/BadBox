package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		faktoriyel := 1
		for j := 1; j <= i; j++ {
			faktoriyel *= j
		}
		fmt.Println(i, "! =", faktoriyel)
	}

}
