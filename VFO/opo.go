package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Lütfen 1.Dönem Sonu Notunuzu Giriniz:")
	scanner.Scan()
	donem1, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Print("Lütfen 2.Dönem Sonu Notunuzu Giriniz:")
	scanner.Scan()
	donem2, _ := strconv.ParseFloat(scanner.Text(), 64)
	donemsonu := (donem1 + donem2) / 2
	fmt.Println("Dönem Sonu Ortalamanız:", donemsonu)
}
