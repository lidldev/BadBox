package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome...")
	fmt.Println("Addition? Extraction? Multiplication? Division?")
	scanner.Scan()
	choose:=strings.ToLower(scanner.Text())
	
	fmt.Println("Please Enter First Number:")
	scanner.Scan()
	number1,err:=strconv.ParseFloat(scanner.Text(),64)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Please Enter Second Number:")
	scanner.Scan()
	number2,err:=strconv.ParseFloat(scanner.Text(),64)
	
	if choose == "addition"{
		fmt.Println(number1+number2)
	}else if choose == "extraciton"{
		fmt.Println(number1-number2)
	}else if choose == "multiplication"{
		fmt.Println(number1*number2)
	}else if choose == "division"{
		fmt.Println(number1/number2)
	}else{
		fmt.Println("Wrong Entry")

	}
	
}