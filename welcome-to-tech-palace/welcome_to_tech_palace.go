package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	borderSymbol := "*"
	finalBorder := strings.Repeat(borderSymbol, numStarsPerLine)
	return finalBorder + "\n" + welcomeMsg + "\n" + finalBorder
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	starsRemoved := strings.ReplaceAll(oldMsg, "*", "")
	wsTrimmed := strings.TrimSpace(starsRemoved)

	return wsTrimmed
}
