package helper
import "strings"

func DataValidation(firstName string, lastName string, email string, userTickets int, remaingTickets int) (bool, bool, bool){

	var isValidName = len(firstName) > 2 && len(lastName) > 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remaingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}