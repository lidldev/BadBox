package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("Çift Sayılar: %d\n", i)
		}
	}
	toplam := 0
	for i := 1; i < 10; i++ {
		toplam += i
	}
	fmt.Println(toplam)
}
