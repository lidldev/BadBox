package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int
	var guess int
	var tries int

	rand.Seed(time.Now().UnixNano())
	num = rand.Intn(100) + 1

	for guess != num {
		fmt.Print("Guess Number Between 1-100: ")
		fmt.Scan(&guess)
		tries++

		if guess > num {
			fmt.Println("Too High")
		} else if guess < num {
			fmt.Println("Too Low")
		} else {
			fmt.Println("Good Job!!!!")
			fmt.Println("Your Tries", tries)
		}
	}
}
