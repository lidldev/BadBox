package main

import "fmt"

func main() {
	var sayi int
	fmt.Print("Sayı Giriniz:")
	fmt.Scanf("%d", &sayi)
	toplam := 0
	for i := 1; i < sayi; i++ {
		if sayi%i == 0 {
			toplam += i
		}
	}
	if toplam == sayi {
		fmt.Printf("%d Sayısı Mükemmel Sayıdır", sayi)
	} else {
		fmt.Printf("%d Sayısı Mükemmel Sayı Değildir", sayi)
	}
}
