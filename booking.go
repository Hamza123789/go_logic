package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "go conference"
	const conferenceTicket = 50
	var remainingTickets = 50

	fmt.Printf("conference name is %T\n", conferenceName)

	fmt.Printf("welcome to %v booking application ", conferenceName)
	fmt.Printf("we have total of %v tickets and %v remaining\n", conferenceTicket, remainingTickets)

	fmt.Println("get you tickets here to attend")

	for i := 0; i < 3; i++ {
		var bookings = []string{}

		var userName string
		var lastName string

		var userTickets int
		var email string

		fmt.Println("enter your name:")
		fmt.Scan(&userName)
		fmt.Println("enter your last name:")
		fmt.Scan(&lastName)
		bookings = append(bookings, userName+" "+lastName)

		fmt.Println("enter your email:")
		fmt.Scan(&email)

		fmt.Println("enter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets -= userTickets

		fmt.Printf("user %v %v booked %v tickets. and will recive them on %v the remaining tickets are %v and the bookings are %v", userName, lastName, userTickets, email, remainingTickets, bookings)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names %v\n", firstNames)

	}

}
