package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Vize Notunuzu Giriniz:")
	scanner.Scan()
	vizenot, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Print("Final Notunuzu Giriniz:")
	finalnot, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	ortalama := (vizenot * 0.4) + (finalnot * 0.6)
	fmt.Println("Ortalamanız:", ortalama)

	if ortalama >= 85 && ortalama <= 100 {
		fmt.Println("AA")
	} else if ortalama >= 50 && ortalama <= 84 {
		fmt.Println("BB")
	} else {
		fmt.Println("Mal")
	}

}
