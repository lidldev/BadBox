package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var flightName = "İstanbul-Barcelona Flight"
	var remaining = 50
	var ticketPrice = 30
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome To Flight Ticket Book App")
	fmt.Println("Please Write Your Name")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("How Much Ticket Do You Need")
	scanner.Scan()
	pieceOfTicket, _ := strconv.ParseFloat(scanner.Text(), 64)
	cost := int(pieceOfTicket) * ticketPrice
	whatWeHave := remaining - int(pieceOfTicket)
	fmt.Printf("%v got %v ticket,it costs %v,for %v,after purchase %v tickets left ", name, pieceOfTicket, cost, flightName, whatWeHave)

}
