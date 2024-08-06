package main

import (
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	scores := map[string]int{
		"Math":    80,
		"Science": 90,
		"History": 70,
	}

	expected := 80.0
	actual := calculateAverage(scores)

	if actual != expected {
		t.Errorf("Expected %.2f but got %.2f", expected, actual)
	}
}

func TestGetGrade(t *testing.T) {
	tests := []struct {
		score    int
		expected string
	}{
		{95, "A"},
		{85, "A"},
		{75, "A"},
		{65, "B"},
		{55, "C"},
		{45, "D"},
		{35, "F"},
	}

	for _, test := range tests {
		actual := getGrade(test.score)
		if actual != test.expected {
			t.Errorf("For score %d, expected grade %s but got %s", test.score, test.expected, actual)
		}
	}
}
