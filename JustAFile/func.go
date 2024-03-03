package main

import "fmt"

func ucislem(x, y int) (int, int, int) {
	return x + y, x - y, x * y
}

func maxmin(x, y int) (int, int) {
	if x > y {
		return x, y
	} else {
		return x, y
	}

}

func main() {
	sayi1, sayi2, sayi3 := ucislem(75, 98)
	fmt.Printf("Sonuç: %d\n", sayi1)
	fmt.Printf("Sonuç: %d\n", sayi2)
	fmt.Printf("Sonuç: %d\n", sayi3)
	sayi, sayi4 := maxmin(983, 786)
	fmt.Printf("Max=%d Min=%d", sayi, sayi4)
}
