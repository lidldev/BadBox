package main

import "fmt"

func main() {
	plaka := make(map[int]string)
	plaka[44] = "Malatya"
	plaka[35] = "İzmir"
	plaka[01] = "Adana"
	fmt.Println(plaka)

	unvanlar := make(map[string]string)
	unvanlar["DR"] = "Doktor"
	unvanlar["ORD"] = "Ordinaryus"
	unvanlar["PRF"] = "Profesör"
	unvanlar["DOÇ"] = "Doçent"
	fmt.Println(unvanlar)
	delete(unvanlar, "DR")
	fmt.Println(unvanlar)
	for key, value := range unvanlar {
		fmt.Printf("Kısaltması=%v\n Tam Hali=%v\n", key, value)
	}
}
