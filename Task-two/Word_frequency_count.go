package main
import (
    "fmt"
    "strings"
)

func WordFrequencyCount(text string) map[string]int {
    words := strings.Fields(text)
    wordFrequency := map[string]int{}
    for _, word := range words {
        wordFrequency[word]++
    }

    return wordFrequency
}

func main() {
    text := "example text with some example words"
    frequencies := WordFrequencyCount(text)
    fmt.Println(frequencies)
	
}