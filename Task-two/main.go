package main

import (
	"fmt"
)
func main() {
    text := "example text with some example words"
    frequencies := WordFrequencyCount(text)
    fmt.Println(frequencies)

	text = "racecar"
	fmt.Println(PalindromeCheck(text))
	
}
