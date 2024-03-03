package main

import "fmt"

func fibore(n int) int {
	if n <= 1 {
		return n
	}
	return fibore(n-1) + fibore(n-2)
}

func main() {
	for i := 0; i <= 20; i++ {
		sonuc := fibore(i)
		fmt.Printf("Fibonacci:%d = %d\n", i, sonuc)
	}

}
