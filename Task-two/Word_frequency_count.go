package main

import (
	"strings"
	"unicode"
)


// Function to count the frequency of each word in a string
func WordFrequency(s string) map[string]int {
	// Create a map to store the frequency of each word
	frequencyMap := make(map[string]int)

	// Convert the string to lowercase
	s = strings.ToLower(s)

	// Function to remove punctuation from a string
	removePunctuation := func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}

	// Remove punctuation from the string
	s = strings.Map(removePunctuation, s)

	// Split the string into words
	words := strings.Fields(s)

	// Count the frequency of each word
	for _, word := range words {
		frequencyMap[word]++
	}

	return frequencyMap
}


