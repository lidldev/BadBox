package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for true {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Lütfen Katına Bakmak İstediğiniz Sayıyı Yazınız:")
		fmt.Println("1'in Katları\n" +
			"2'nin Katları\n" +
			"3'ün Katları\n" +
			"4'ün Katları\n" +
			"5'in Katları\n" +
			"6'nın Katları\n" +
			"7'nin Katları\n" +
			"8'in Katları\n" +
			"9'un Katları")
		scanner.Scan()

		secim := scanner.Text()
		if secim == "1" {
			for i := 0; i <= 1000; i++ {
				if i%1 == 0 {
					fmt.Println(i)
				}
			}
		}
		if secim == "2" {
			for x := 0; x <= 1000; x++ {
				if x%2 == 0 {
					fmt.Println(x)
				}
			}

		}
		if secim == "3" {
			for y := 0; y <= 1000; y++ {
				if y%3 == 0 {
					fmt.Println(y)
				}
			}
		}
		if secim == "4" {
			for z := 0; z <= 1000; z++ {
				if z%4 == 0 {
					fmt.Println(z)
				}
			}
		}
		if secim == "5" {
			for q := 0; q <= 1000; q++ {
				if q%5 == 0 {
					fmt.Println(q)
				}
			}
		}
		if secim == "6" {
			for q := 0; q <= 1000; q++ {
				if q%6 == 0 {
					fmt.Println(q)
				}
			}
		}
		if secim == "7" {
			for q := 0; q <= 1000; q++ {
				if q%7 == 0 {
					fmt.Println(q)
				}
			}
		}
		if secim == "8" {
			for q := 0; q <= 1000; q++ {
				if q%8 == 0 {
					fmt.Println(q)
				}
			}
		}
		if secim == "9" {
			for q := 0; q <= 1000; q++ {
				if q%9 == 0 {
					fmt.Println(q)
				}
			}
		}
		if secim == "q" || secim == "Q" {
			break
		}

	}
}
