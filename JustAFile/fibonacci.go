package main

import "fmt"

func fibo(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		gecici := a
		a = b
		b = gecici + a
	}
	return a
}

func main() {
	for c := 0; c <= 20; c++ {
		sonuc := fibo(c)
		fmt.Printf("Fibonacci:%d = %d\n", c, sonuc)
	}

}
