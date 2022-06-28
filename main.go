package main

import (
	"fmt"
	"strings"
)

func main() {
	var applicationName = "Anmol"
	const applicationTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Println("Welcome to our", applicationName, "application")
	fmt.Println("We have total of", applicationTickets, "tickets and", remainingTickets, "are still remaining")
	fmt.Println("Get your tickets to explore our application further")

	for {
		var firstName string
		var lastName string
		var userTickets uint
		var email string
		//ask user for their info and also their required number of tickets
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		bookings = append(bookings, firstName+" "+lastName)
		remainingTickets = remainingTickets - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a email confirmation on %v\n", firstName, lastName, userTickets, email)
		fmt.Println("We have total of", applicationTickets, "tickets and", remainingTickets, "are still remaining")

		bookedFirstNames := []string{}
		for _, booking := range bookings {
			bookedFirstNames = append(bookedFirstNames, strings.Fields(booking)[0])
		}
		fmt.Printf("These are our all the bookings:%v\n", bookedFirstNames)
	}
}
