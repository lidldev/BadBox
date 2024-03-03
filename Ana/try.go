package main

import "fmt"

func main() {

	for x := 0; x <= 100; x++ {
		if x%2 == 0 {
			fmt.Print(x, " ")
			x++
		}
	}

	for y := 0; y < 100; y++ {
		if y%3 == 0 {
			fmt.Print(y, " ")
			y++
		}
	}

	for z := 0; z <= 100; z++ {
		if z%5 == 0 {
			fmt.Print(z, " ")
			z++
		}
	}

	for w := 0; w <= 100; w++ {
		if w%7 == 0 {
			fmt.Print(w, " ")
			w++
		}
	}

	for q := 0; q <= 100; q++ {
		if q%9 == 0 {
			fmt.Print(q, " ")
			q++
		}
	}

}
