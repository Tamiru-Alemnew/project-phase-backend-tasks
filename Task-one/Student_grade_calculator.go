package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Student struct {
    Name         string
    SubjectScores map[string]int
    AverageGrade  float64
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("What is your name?")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    fmt.Println("Please provide the number of subjects you have taken")
    numberOfSubjectsStr, _ := reader.ReadString('\n')
    numberOfSubjects, err := strconv.Atoi(strings.TrimSpace(numberOfSubjectsStr))
    if err != nil || numberOfSubjects <= 0 {
        fmt.Println("Invalid number of subjects. Please enter a positive integer.")
        return
    }

    subjectToScore := make(map[string]int)

    for i := 0; i < numberOfSubjects; i++ {
        fmt.Println("Enter the subject name")
        subjectName, _ := reader.ReadString('\n')
        subjectName = strings.TrimSpace(subjectName)

        fmt.Println("Enter the subject score (0-100)")
        subjectScoreStr, _ := reader.ReadString('\n')
        subjectScore, err := strconv.Atoi(strings.TrimSpace(subjectScoreStr))
        if err != nil || subjectScore < 0 || subjectScore > 100 {
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