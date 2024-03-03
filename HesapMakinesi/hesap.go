package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Lütfen Seçim Yapınız")
	fmt.Println("1-Toplama\n2-Çıkarma\n3-Çarpma\n4-Bölme\n5-Mod Alma")
	scanner.Scan()
	secim := scanner.Text()
	fmt.Print("1. Sayıyı Giriniz")
	scanner.Scan()
	sayi1, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Print("2. Sayıyı Giriniz")
	scanner.Scan()
	sayi2, _ := strconv.ParseFloat(scanner.Text(), 64)

	if secim == "1" {
		fmt.Printf("Toplama %f", sayi1+sayi2)
	} else if secim == "2" {
		fmt.Printf("Çıkarma %f", sayi1-sayi2)
	} else if secim == "3" {
		fmt.Printf("Çarpma %f", sayi1*sayi2)
	} else if secim == "4" {
		fmt.Printf("Bölme %f", sayi1/sayi2)
	} else if secim == "5" {
		modAlma := int(sayi1) % int(sayi2)
		fmt.Printf("Mod Alma %d", modAlma)
	} else {
		fmt.Println("Birşey Yazmadan İşlem Yapmazsın")
	}
}
