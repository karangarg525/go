package main

import (
	"fmt"
	"firstapp/greet"
	"strconv"
)

func main() {
	var conferenceName string = "Go Conference"
	var conferenceTickets int = 50
	var remainingTickets int = 50
	var bookings = make([]map[string]string, 0)

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

func bookTicket(firstName string, lastName string, noOfTickets int, email string, remainingTickets int, bookings []map[string]string) ([]map[string]string, int) {
	remainingTickets = remainingTickets - noOfTickets
	fmt.Printf("Hi, %v %v, You have booked %v tickets. You will get the confirmation on %v.\nNow there are %v tickets left\n", firstName, lastName, noOfTickets, email, remainingTickets)

	// Creating a Map for a user
	var userData = make(map[string]string) 
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["noOfTickets"] = strconv.Itoa(noOfTickets)

	return append(bookings, userData), remainingTickets
}

func printFirstNames(bookings []map[string]string) {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	fmt.Printf("Bookings: %v\n", firstNames)
	fmt.Println(bookings)
}
