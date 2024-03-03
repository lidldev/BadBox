package main

import "fmt"

func main() {
	sayi := 20
	switch sayi {
	case 5:
		fmt.Println("Sayı 5'e Eşittir")
	case 10:
		fmt.Println("Sayı 10'a Eşittir")
	case 15:
		fmt.Println("Sayı 15'e Eşittir")
	default:
		fmt.Println("Yalan Yapma Yalan Yapmaa!!!")
	}
}
