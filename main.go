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

	greetUsers(applicationName, applicationTickets, remainingTickets)

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

		//Data Validation
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a email confirmation on %v\n", firstName, lastName, userTickets, email)
			fmt.Println("We have total of", applicationTickets, "tickets and", remainingTickets, "are still remaining")

			bookedFirstNames := getFirstNames(bookings)
			fmt.Printf("The first name of all the bookings: %v\n", bookedFirstNames)

			if remainingTickets == 0 {
				fmt.Println("Tickets for our application is booked out. Come back in next slot")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short!")
			}
			if !isValidEmail {
				fmt.Println("e-mail address you entered doesn't contain @ symbol!")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of ticket you entered is either less than 0 or greater than remaining tickets!")
			}
		}
	}
}

func greetUsers(applicationName string, applicationTickets uint, remainingTickets uint) {
	fmt.Println("Welcome to our", applicationName, "application")
	fmt.Println("We have total of", applicationTickets, "tickets and", remainingTickets, "are still remaining")
	fmt.Println("Get your tickets to explore our application further")
}

func getFirstNames(bookings []string) []string {
	bookedFirstNames := []string{}
	for _, booking := range bookings {
		bookedFirstNames = append(bookedFirstNames, strings.Fields(booking)[0])
	}
	return bookedFirstNames
}
