package main

import "fmt"

func main() {
	var conferenceName = "Go Conference" //variable declaration
	const conferenceTickets = 50 //constant decalration
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName,  "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "tickets are left.")
	fmt.Println("Get your tickets here to attend")

}