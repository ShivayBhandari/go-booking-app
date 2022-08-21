package main

import (
	"fmt"
	"strings"
	"go-booking-app/helper"
)

var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {
		
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			
			bookTicket(userTickets, firstName, lastName, email)

			//call function print first names
			printFirstNames()

			if remainingTickets == 0 {
				
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName{
				fmt.Println("first name or last name is too short")
			}
			if !isValidEmail{
				fmt.Println("email does not contain @")
			}
			if !isValidTicketNumber{
				fmt.Println("number of ticket is invalid")
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames(){
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	fmt.Printf("The first names of bookings are: %v\n", firstNames)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}