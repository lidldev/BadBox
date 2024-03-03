package main

import (
	"log"
	"os"
)

func main() {

	os.Create("data.txt")
	dosya, err := os.OpenFile("data.txt", os.O_WRONLY|os.O_APPEND, 0666)
	defer dosya.Close()
	if err != nil {
		log.Fatal(err)
	}
	dosya.WriteString("Ankara\nİstanbul\nBursa\nGiresun")

}
