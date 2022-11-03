package greet

import (
	"fmt"
)

func GreetUser(conferenceName string, conferenceTickets int, remainingTickets int) {
	fmt.Printf("Welcome to %v ticket booking application\n", conferenceName)
	fmt.Printf("We have total %v, tickets and %v, are available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n\n")
}
