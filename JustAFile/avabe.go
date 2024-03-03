package main

import "fmt"

type calisan struct {
	isim      string
	yas       int
	maas      int
	isHaveKid bool
}

type mudur struct {
	calisan
	unvan string
}

func main() {
	a1 := calisan{
		isim:      "Ahmet Affan",
		yas:       38,
		maas:      11402,
		isHaveKid: false,
	}
	fmt.Println(a1)

	y1 := mudur{
		calisan: calisan{
			isim:      "Münür",
			yas:       51,
			maas:      51394,
			isHaveKid: true,
		},
		unvan: "Müdür Bey",
	}
	fmt.Println(y1)

	kurucu := struct {
		isim    string
		sermaye int
	}{isim: "Mehmet Şimşek", sermaye: 1453000}

	fmt.Println(kurucu)

}
