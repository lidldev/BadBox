package main

import "fmt"

func main() {
	var a []int = []int{11, 22, 33, 44, 55}
	fmt.Println(a)
	a = append(a, 66)
	fmt.Println(a)
	a = append(a, 77)
	fmt.Println(a)
	a = append(a, 88)
	fmt.Println(a)

	var sehirler = []string{"Malatya", "Bursa", "Antalya"}
	sehirler = append(sehirler, "İstanbul")
	sehirler = append(sehirler, "İzmir")
	fmt.Println(sehirler)

	sayilar := make([]int, 5)
	sayilar[0] = 3
	sayilar[1] = 6
	sayilar[2] = 9
	sayilar[3] = 12
	sayilar[4] = 15
	fmt.Println(sayilar)

}
