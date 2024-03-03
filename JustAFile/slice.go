package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ulkeler := make([]string, 5)

	for i := 0; i < 6; i++ {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Ülke Giriniz:")
		scanner.Scan()
		ulkegiris := scanner.Text()
		ulkeler = append(ulkeler, ulkegiris)

	}
	fmt.Println(ulkeler)
	for index, eleman := range ulkeler {
		if index < 5 {
			continue
		}
		fmt.Printf("Girdiğiniz Ülke:%d. Ülke:%s\n", index-4, eleman)
	}
}
