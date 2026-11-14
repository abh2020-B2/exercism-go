package techpalace

import str "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return ("Welcome to the Tech Palace, " + str.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return str.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + str.Repeat("*", numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return str.TrimSpace(str.ReplaceAll(oldMsg, "*", " "))
}
