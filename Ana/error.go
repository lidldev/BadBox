package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	/*
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Sayı Giriniz:")
		scanner.Scan()
		a, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			//fmt.Println(err.Error())
			log.Fatal("HATA ", err)

		}

		fmt.Println(a)
	*/
	file, err := os.Open("/main.go")
	if err != nil {
		fmt.Println(errors.New("Dosya Yok"))
	}
	fmt.Println(file.Name(), "Başarılı Bir Şekilde Açıldı")

}
