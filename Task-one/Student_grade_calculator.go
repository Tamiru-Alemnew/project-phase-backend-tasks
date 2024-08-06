package main

import (
	"fmt"
)

type Student struct {
	Name         string
	SubjectScores map[string]int
	AverageGrade  float64
}

func main() {
	var name string
	var numberOfSubjects int

	fmt.Println("What is your name?")
	fmt.Scanln(&name)

	fmt.Println("Please provide the number of subjects you have taken")
	fmt.Scanln(&numberOfSubjects)
	if numberOfSubjects <= 0 {
		fmt.Println("Invalid number of subjects. Please enter a positive integer.")
		return
	}

	subjectToScore := make(map[string]int)

	for i := 0; i < numberOfSubjects; i++ {
		var subjectName string
		var subjectScore int

		fmt.Println("Enter the subject name")
		fmt.Scanln(&subjectName)

		fmt.Println("Enter the subject score (0-100)")
		fmt.Scanln(&subjectScore)
		if subjectScore < 0 || subjectScore > 100 {
			fmt.Println("Invalid score. Please enter a score between 0 and 100.")
			i-- // Retry this subject
			continue
		}

		subjectToScore[subjectName] = subjectScore
	}

	student := Student{
		Name:         name,
		SubjectScores: subjectToScore,
		AverageGrade:  calculateAverage(subjectToScore),
	}

	displayResults(student)
}

func calculateAverage(scores map[string]int) float64 {
	var sum int
	for _, score := range scores {
		sum += score
	}
	return float64(sum) / float64(len(scores))
}

func getGrade(score int) string {
	switch {
	case score >= 70:
		return "A"
	case score >= 60:
		return "B"
	case score >= 50:
		return "C"
	case score >= 40:
		return "D"
	default:
		return "F"
	}
}

func displayResults(student Student) {
	fmt.Printf("\nStudent Name: %s\n", student.Name)
	fmt.Println("Subjects and Scores:")
	for subject, score := range student.SubjectScores {
		fmt.Printf("%s: %d (%s)\n", subject, score, getGrade(score))
	}
	fmt.Printf("Average Grade: %.2f\n", student.AverageGrade)
}
