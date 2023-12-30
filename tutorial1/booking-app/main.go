package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

const conferenceName string = "GO Conference"
const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers(conferenceName, conferenceTickets, remainingTickets)
	// var bookings []string
	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidUserTickets {
			bookTicket(userTickets, firstName, lastName, email, conferenceName)
			// call function print firstName
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			//var noTicketsRemaining bool = remainingTickets == 0
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidUserTickets {
				fmt.Println("number of tickets you entered is invalid.")
			}
			fmt.Println("Your input data is invalid, Please try again.")
			continue
		}
	}
}

func greetUsers(confName string, confTicket int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTicket, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	// create a map for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
