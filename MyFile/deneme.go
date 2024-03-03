package main

import "fmt"

type ogrenci struct {
	isim string
	yas  int
	not  []int
}

func main() {
	var o1 ogrenci

	o1.isim = "Ahmet Affan Ebcioğlu"
	o1.yas = 14
	o1.not = []int{100, 90, 97}
	fmt.Println(o1)

}
