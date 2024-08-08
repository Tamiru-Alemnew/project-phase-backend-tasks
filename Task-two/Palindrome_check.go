package main
import (
	"strings"
	"unicode"
)

// Function to check if a string is a palindrome
func IsPalindrome(s string) bool {
	// Helper function to remove non-alphanumeric characters and convert to lowercase
	normalize := func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return -1
	}

	// Normalize the input string
	normalizedStr := strings.Map(normalize, s)

	// Check if the normalized string is a palindrome
	n := len(normalizedStr)
	for i := 0; i < n/2; i++ {
		if normalizedStr[i] != normalizedStr[n-1-i] {
			return false
		}
	}
	return true
}


