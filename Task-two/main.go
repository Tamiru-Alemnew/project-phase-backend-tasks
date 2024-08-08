package main

import (
	"fmt"
)
func main() {
    text := "example text with some example words"
    frequencies := WordFrequency(text)
    fmt.Println(frequencies)

	text = "racecar"
	fmt.Println(IsPalindrome(text))
	
}
