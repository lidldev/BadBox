package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Your Age")
	scanner.Scan()
	iage, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i <= 4; i++ {
		if (iage >= 1) && (iage <= 18) {
			pl("Best Age")
		} else if iage == 49 && iage == 51{
			pl("Best Age <3")
		} else {
			pl("Hmm")
		}
	}
}
