package main

import "fmt"

func main() {
	var user []string

	fmt.Println("Welcome To Ticket App...")
	fmt.Println("Enter your name,last name,E-mail and how many tickets do you want")

	var name string
	var lname string
	var email string
	const allTickets int = 60
	var avaliableTickets uint = 60
	var userTicket uint

	for true {

		fmt.Println("Enter Your Name:")
		fmt.Scan(&name)
		fmt.Println("Enter Your Last Name:")
		fmt.Scan(&lname)
		fmt.Println("Enter Your Email Adress:")
		fmt.Scan(&email)
		fmt.Println("Enter How Many Tickets Do You Want")
		fmt.Scan(&userTicket)

		user = append(user, name)

		avaliableTickets = avaliableTickets - userTicket

		fmt.Printf("You buy %v tickets and %v tickets are left\n", userTicket, avaliableTickets)
		fmt.Println(user)

		if avaliableTickets < 0 {
			fmt.Println("Out Of Ticket:")
			break
		}

	}

}

