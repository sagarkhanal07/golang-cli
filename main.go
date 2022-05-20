package main

//Core package for different functionalites e.g Print
import (
	"fmt"
)

//Entrypoint for GO Execution
func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 12
	var remainingTickets uint = 12
	bookings := []string{}

	//space automatically added for variables
	fmt.Printf("Welcome to %q booking application\n", conferenceName)
	fmt.Printf("Total tickets available: %v\n", conferenceTickets)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
	fmt.Println("Get your tickets here 🎫 to attend the conference 🗣")

	for remainingTickets > 0 {
		var userName string
		var userTickets uint
		var email string

		//ask user for input
		fmt.Print("Enter your name: ")
		fmt.Scan(&userName)

		//ask input for no of email
		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		//ask input for no of tickets
		fmt.Print("Enter no of tickets: ")
		fmt.Scan(&userTickets)
		bookings = append(bookings, userName)

		//ticket status
		if userTickets <= conferenceTickets {

			fmt.Printf("%v tickets booked for %v\n", userTickets, userName)
			remainingTickets = conferenceTickets - userTickets
			fmt.Printf("Remaining tickets: %v\n", remainingTickets)
			fmt.Printf("Thank you for buying tickets. Tickets has been sent to %s\n", email)

			firstNames := []string{}
			for _, booking := range bookings {
				firstNames = append(firstNames, booking[:1])
			}
			fmt.Println("Bookings:", firstNames)

			// noTicketsRemaining := remainingTickets == 0

			if remainingTickets == 0 {
				fmt.Println("Sorry, we are sold out. Please try again later.")
				break
			}
		} else {
			fmt.Println("We only have ", conferenceTickets, " tickets available")
			// break // End booking ticket
			// continue // Continue booking ticket
		}

	}

}
