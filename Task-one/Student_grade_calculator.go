package main

import (
    "fmt"
)


func main(){
	var name string
	var number_of_suject int
	
	fmt.Println("What is your name?")
	fmt.Scanln(&name)
	fmt.Println("Please provide the number of subjects you have taken")
	fmt.Scanln(&number_of_suject)
	fmt.Println(name , number_of_suject)

	subject_to_score := map[string]int{}

	for i:= 0 ; i < number_of_suject;i++ {
		var subject_name string
		var subject_score int

		fmt.Println("Enter the subject name")
		fmt.Scanln(&subject_name)
		fmt.Println("Enter the subject score")
		fmt.Scanln(&subject_score)
		subject_to_score[subject_name] = subject_score

	}

	fmt.Println("This is the subjects and their scores")

	for key, value := range subject_to_score{
		var grade string
		if value >= 70{
			grade = "A"
		}else if value >= 60{
			grade = "B"
		}else if value >= 50{
			grade = "C"
		}else if value >= 40{
			grade = "D"
		}else{
			grade = "F"
		}
		fmt.Println(key, value, grade)
	}

	var average_score int

	for _, value := range subject_to_score{
		average_score += value
	}

	average_score = average_score/number_of_suject

	fmt.Println("The average score is ", average_score)
}