package main

import "fmt"

type adres struct {
	sehir string
	ulke  string
}

type kisi struct {
	isim      string
	yas       int
	kisiadres adres
}

func main() {
	k1 := kisi{
		isim: "Ahmet Affan",
		yas:  14,
		kisiadres: adres{
			sehir: "Bursa",
			ulke:  "Türkiye",
		},
	}
	fmt.Println("İsim:", k1.isim)
	fmt.Println("Yaş:", k1.yas)
	fmt.Println("Şehir:", k1.kisiadres.sehir)
	fmt.Println("Ülke:", k1.kisiadres.ulke)
}
