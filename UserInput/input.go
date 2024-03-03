package main

import "fmt"

func main() {
	var isim string
	var yas int

	fmt.Print("Adınızı Yazınız:")
	fmt.Scan(&isim)

	fmt.Print("Yaşınızı Yazınız:")
	fmt.Scan(&yas)

	fmt.Printf("Adınız:%s Yaşınız:%d", isim, yas)
}
