package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	pl("What is Your Name")
	scanner.Scan()
	name := scanner.Text()
	pl("How Old Are You")
	scanner.Scan()
	age, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("Your Name: %s\nYour Age:%d", name, age)

}
