// Package raindrops does a fizz buzz like thing
package raindrops

import (
	"strconv"
	"strings"
)

// Convert converts a number to a string with some text in if it matches the rules
func Convert(number int) string {
	var sb strings.Builder

	if number%3 == 0 {
		sb.WriteString("Pling")
	}

	if number%5 == 0 {
		sb.WriteString("Plang")
	}

	if number%7 == 0 {
		sb.WriteString("Plong")
	}

	result := sb.String()

	if result == "" {
		result = strconv.Itoa(number)
	}

	return result
}