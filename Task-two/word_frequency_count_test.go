package main

import (
    "reflect"
    "testing"
)

func TestWordFrequencyCount(t *testing.T) {
    text := "example text with some example words example"
    expected := map[string]int{
        "example": 3,
        "text":    1,
        "with":    1,
        "some":    1,
        "words":   1,
    }

    result := WordFrequency(text)

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, but got %v", expected, result)
    }
}