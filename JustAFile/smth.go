package main

import "fmt"

func main() {
	var dizi1 [2][3]int
	var dizi2 [2][3]int

	var dizisonuc [2][3]int
	fmt.Println("1.Dizinin Elemnalarını Girin:\n")
	for k, r := range dizi1 {
		for l := range r {
			fmt.Scan(&dizi1[k][l])
		}
	}
	fmt.Println("2.Dizinin Elemanlarını Girin:\n")
	for p, c := range dizi2 {
		for f := range c {
			fmt.Scan(&dizi2[p][f])
		}
	}
	fmt.Println("Dizilerin Çarpımı Sonucu:\n")
	for c, d := range dizisonuc {
		for e := range d {
			dizisonuc[c][e] = dizi1[c][e] * dizi2[c][e]
			fmt.Println(dizisonuc[c][e])
		}
	}
	fmt.Println()

}
