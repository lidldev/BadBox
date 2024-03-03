package main

import (
	"fmt"
	"math"
)

const PI = 3.14

func daire(yaricap float64) (alan float64, cevre float64) {
	alan = PI * (math.Pow(yaricap, 2.0))
	cevre = PI * 2 * yaricap
	return
}

func main() {
	alan, cevre := daire(5.0)
	fmt.Printf("Dairenin Alanı: %f\nDairenin Çevresi: %f", alan, cevre)

}
