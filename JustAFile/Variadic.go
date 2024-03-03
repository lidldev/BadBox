package main

import "fmt"

func recursive() {
	fmt.Println("Merhaba")
	recursive()

}

func faktoriyel(sayi int) int {
	if sayi <= 1 {
		return 1
	}
	return sayi * faktoriyel(sayi-1)
}

func main() {
	//recursive()
	sayi := 50
	fmt.Printf("%d Sayısının Faktoriyeli: %d", sayi, faktoriyel(sayi))
}
