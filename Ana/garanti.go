package main

import (
	"fmt"
	"time"
)

func Todo(whatTodo string) {
	for i := 0; i <= 100; i++ {
		fmt.Println(i, whatTodo)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go Todo("I Love You")
	Todo("I Am Sorry")

}
