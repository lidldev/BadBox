package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	fmt.Println("Welcome You Are The One Of Lucky Guys :)")
	fmt.Println("If You Enter Our Perfect Number You Win The Prize :)")
	scanner.Scan()
	prize := scanner.Text()

	Computer()
}
func Computer(prize int) int {
	rand.Seed(time.Now().UnixNano())
	c := rand.Intn(100) + 1
	prize = 50
	if c == prize {
		fmt.Println("You Won!!!!!")
	} else {
		fmt.Println("Nice Try :(")
	}
	return prize
}
