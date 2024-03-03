package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for true {
		var sayi int
		scanner := bufio.NewScanner(os.Stdin)
		secim := scanner.Text()
		fmt.Print("Lütfen Bir Sayı Giriniz:")
		scanner.Scan()
		if sayi%2 == 0 {
			fmt.Println("Sayınız Çift Sayıdır")
		} else {
			fmt.Println("Sayınız Çift Sayı Değildir")
		}
		if secim == "q" || secim == "Q" {
			break

		}

	}

}
