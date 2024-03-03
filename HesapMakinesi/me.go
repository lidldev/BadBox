package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Println("1-Toplama\n2-Çıkarma\n3-Çarpma\n4-Bölme")
		scanner.Scan()
		secim := scanner.Text()
		fmt.Print("1.Sayıyı Giriniz:")
		scanner.Scan()
		sayi1, _ := strconv.ParseFloat(scanner.Text(), 64)
		fmt.Print("2.Sayıyı Giriniz")
		scanner.Scan()
		sayi2, _ := strconv.ParseFloat(scanner.Text(), 64)
		if secim == "1" {
			fmt.Printf("%.2f + %.2f = %.2f", sayi1, sayi2, sayi1+sayi2)
		} else if secim == "2" {
			fmt.Printf("%.2f - %.2f = %.2f", sayi1, sayi2, sayi1-sayi2)
		} else if secim == "3" {
			fmt.Printf("%.2f * %.2f = %.2f", sayi1, sayi2, sayi1*sayi2)
		} else if secim == "4" {
			fmt.Printf("%.2f / %.2f = %.2f", sayi1, sayi2, sayi1/sayi2)
		} else {
			fmt.Println("Hatalı Giriş")
		}
	}
}
