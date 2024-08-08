package main
import (
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

