package main

import (
	"fmt"
	"firstapp/greet"
)

type userData struct {
	firstName string
	lastName string
	email string
	noOfTickets int
}

func main() {
	var conferenceName string = "Go Conference"
	var conferenceTickets int = 50
	var remainingTickets int = 50
	// var bookings = make([]map[string]string, 0)
	var bookings = make([]userData, 0)



	greet.GreetUser(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, noOfTickets := getUserDetails()

		if noOfTickets > remainingTickets {
			fmt.Printf("You cannot book %v tickets, there are only %v ticket left\n\n", noOfTickets, remainingTickets)
			continue
		}

		bookings, remainingTickets = bookTicket(firstName, lastName, noOfTickets, email, remainingTickets, bookings)
		
		printFirstNames(bookings)
		
		if remainingTickets == 0 {
			fmt.Println("We are sold out!!!")
			break
		}

	}

}



func getUserDetails() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var noOfTickets int

	fmt.Printf("Enter your First name: ")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your Last name: ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your Email: ")
	fmt.Scan(&email)
	fmt.Printf("Enter the no. of tickets you want to book: ")
	fmt.Scan(&noOfTickets)

	return firstName, lastName, email, noOfTickets
}

func bookTicket(firstName string, lastName string, noOfTickets int, email string, remainingTickets int, bookings []userData) ([]userData, int) {
	remainingTickets = remainingTickets - noOfTickets
	fmt.Printf("Hi, %v %v, You have booked %v tickets. You will get the confirmation on %v.\nNow there are %v tickets left\n", firstName, lastName, noOfTickets, email, remainingTickets)

	// Creating a Struct for a user
	var userData = userData { 
		firstName: firstName,
		lastName: lastName,
		email: email,
		noOfTickets: noOfTickets,
	}

	return append(bookings, userData), remainingTickets
}

func printFirstNames(bookings []userData) {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("Bookings: %v\n", firstNames)
	fmt.Println(bookings)
}
