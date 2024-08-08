package main

import (
	"reflect"
	"testing"
)

func TestPalindromeCheck(t *testing.T) {
	text := "madam" 
	expected := true 
		
	result := IsPalindrome(text)

	if !reflect.DeepEqual(result, expected) {

		t.Errorf("Expected %v, but got %v", expected, result)
	}
}