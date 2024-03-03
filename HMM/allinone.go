package main

import (
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hello World\n")
	time.Sleep(3 * time.Second)
	file.WriteString("Im Ahmet")
	defer file.Close()

	file1, err := os.Create("veri.txt")
	file1.WriteString("Merhaba Dünya\n")
	time.Sleep(3 * time.Second)
	file1.WriteString("Ben Ahmet")
	defer file1.Close()

}
