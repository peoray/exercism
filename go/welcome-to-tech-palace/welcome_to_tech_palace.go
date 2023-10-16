// Package techpalace implements an appliance store
package techpalace

import (
    "strings"
    "fmt"
    )

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    repeatStars := strings.Repeat("*", numStarsPerLine)
    return fmt.Sprintf("%s\n%s\n%s", repeatStars, welcomeMsg, repeatStars)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
