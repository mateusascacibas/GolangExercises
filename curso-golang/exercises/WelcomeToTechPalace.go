package Main

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	i := 0
	j := 0
	var response string
	for i < numStarsPerLine {
		response += "*"
		i++
	}
	response += "\n" + welcomeMsg + "\n"
	for j < numStarsPerLine {
		response += "*"
		j++
	}
	return response
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	newMsg := strings.ReplaceAll(oldMsg, "*", "")
	newMsg = strings.TrimSpace(" \t\n" + newMsg + "\n\t\r\n")
	return newMsg
}
