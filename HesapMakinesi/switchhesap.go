package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("1-Toplama\n2-Çıkarma\n3-Çarpma\n4-Bölme\nLütfen İşlem Seçiniz")
	scanner.Scan()
	secim := scanner.Text()
	fmt.Print("Lütfen 1.Sayıyı Giriniz:")
	scanner.Scan()
	sayi1, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Print("Lütfen 2.Sayıyı Giriniz")
	scanner.Scan()
	sayi2, _ := strconv.ParseFloat(scanner.Text(), 64)
	switch secim {
	case "1":
		fmt.Printf("Toplama %f", sayi1+sayi2)
	case "2":
		fmt.Printf("Çıkarma %f", sayi1-sayi2)
	case "3":
		fmt.Printf("Çarpma %f", sayi1*sayi2)
	case "4":
		fmt.Printf("Bölme %f", sayi1/sayi2)
	default:
		fmt.Printf("Sayı Yazmadan İşlem Yapamazsınız")
	}
}
