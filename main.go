package main

import (
	"fmt"
	"go-experiment/helper"
	"sync"
	"time"
)

//Package Level Variables
var applicationName = "Anmol"
var remainingTickets uint = 50

//List of maps
//var bookings = make([]map[string]string, 0)

var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

const applicationTickets = 50

var wg = sync.WaitGroup{}

//Here our program starts
func main() {

	greetUsers()

	/*for {
		//Get user input
		firstName, lastName, userTickets, email := getUserInput()

		//Data Validation
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			go printTicket(userTickets, firstName, lastName, email)
			bookedFirstNames := getFirstNames()
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
	}*/
	//Get user input
	firstName, lastName, userTickets, email := getUserInput()

	//Data Validation
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go printTicket(userTickets, firstName, lastName, email)

		bookedFirstNames := getFirstNames()
		fmt.Printf("The first name of all the bookings: %v\n", bookedFirstNames)

		if remainingTickets == 0 {
			fmt.Println("Tickets for our application is booked out. Come back in next slot")
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
	wg.Wait()
}

func greetUsers() {
	fmt.Println("Welcome to our", applicationName, "application")
	fmt.Println("We have total of", applicationTickets, "tickets and", remainingTickets, "are still remaining")
	fmt.Println("Get your tickets to explore our application further")
}

func getFirstNames() []string {
	bookedFirstNames := []string{}
	for _, booking := range bookings {
		//bookedFirstNames = append(bookedFirstNames, booking["firstName"])
		bookedFirstNames = append(bookedFirstNames, booking.firstName)
	}
	return bookedFirstNames
}

func getUserInput() (string, string, uint, string) {
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
	return firstName, lastName, userTickets, email
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//Create an empty map
	/*var userData = make(map[string]string)

	//Insert values into map
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)*/

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//Append this map to bookings list
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings so far %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a email confirmation on %v\n", firstName, lastName, userTickets, email)
	fmt.Println("We have total of", applicationTickets, "tickets and", remainingTickets, "are still remaining")
}

func printTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	fmt.Println("################################################################################")
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("Sending %v to Email address %v\n", ticket, email)
	fmt.Println("################################################################################")
	wg.Done()
}
