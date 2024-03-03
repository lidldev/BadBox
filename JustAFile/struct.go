package main

import "fmt"

type ogrenci struct {
	isim      string
	soyisim   string
	yas       int
	notlar    []int
	ismezunmu bool
}

func main() {
	var o1 ogrenci

	o1.isim = "Ahmet Affan"
	o1.soyisim = "Ebcioğlu"
	o1.yas = 14
	o1.notlar = []int{98, 97, 99}
	o1.ismezunmu = false
	fmt.Println(o1)

	o2 := ogrenci{
		isim:      "Özgür",
		soyisim:   "Demirtaş",
		yas:       45,
		notlar:    []int{100, 100, 100},
		ismezunmu: true,
	}

	fmt.Println(o2)

}
