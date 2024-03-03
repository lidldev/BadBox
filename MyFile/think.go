package main

import "fmt"

type ogrenci struct {
	isim    string
	yas     int
	notlar  []int
	gectimi bool
}

func (o ogrenci) adgoster() string {
	return o.isim
}

func (ogr ogrenci) blggoster() {
	fmt.Println("Öğrenci Adı:", ogr.isim)
	fmt.Println("Öğrenci Yaşı:", ogr.yas)
	fmt.Println("Öğrenci Notları:", ogr.notlar)
	fmt.Println("Öğrenci Geçtimi:", ogr.gectimi)

}

func (not ogrenci) ortalama() float64 {
	toplam := 0
	for _, i := range not.notlar {
		toplam += i

	}
	return float64(toplam) / float64(len(not.notlar))
}

func (gec ogrenci) geciyomu() bool {
	if gec.ortalama() <= 70 {
		return false
	}
	return true
}

func main() {
	o1 := ogrenci{"AAE", 14, []int{97, 91, 100}, true}
	fmt.Println(o1.adgoster())
	o1.blggoster()
	fmt.Println(o1.ortalama())
	fmt.Println(o1.geciyomu())
}
