package main

import (
	"fmt"
	"time"
)

type ustunf float64
type Aylar int

const (
	Ocak Aylar = 1 + iota
	Subat
	Mart
	Nisan
	Mayis
	Haziran
	Temmuz
	Agustos
	Eylul
	Ekim
	Kasim
	Aralik
)

func (y ustunf) beslecarp() ustunf {
	return y * 5.2
}

func main() {
	var o1 ustunf = 5.757
	fmt.Println(o1)
	var o2 float64 = 9.81
	fmt.Println(o1 + ustunf(o2))
	y1 := ustunf(50.5)
	fmt.Println(y1.beslecarp())

	_, ay, _ := time.Now().Date()
	fmt.Println(ay)
	fmt.Println(Aylar(ay))
}
