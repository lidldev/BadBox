package main

import "fmt"

func main() {
	any := func(toplam int) int {
		return toplam / 3 * 4 * 5
	}(10)

	fmt.Println(any)
}
