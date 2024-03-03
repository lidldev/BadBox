package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Yatay Kenarı Giriniz:")
	scanner.Scan()
	yatay, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Println("Dikey Kenarı Giriniz:")
	scanner.Scan()
	dikey, _ := strconv.ParseFloat(scanner.Text(), 64)
	z := math.Hypot(yatay, dikey)
	fmt.Println(z)
}
