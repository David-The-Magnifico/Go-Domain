package main

import (
	"fmt"
)

func main() {
	questions := []string{
		"1. What is the primary greenhouse gas emitted by human activities?",
		"2. What is the main cause of the melting of polar ice caps?",
		"3. What is the consensus among scientists regarding the primary cause of global warming?",
		"4. Which of the following is a potential consequence of global warming?",
		"5. What are some proposed solutions to mitigate global warming?",
	}

	options := [][]string{
		{"1. Carbon Dioxide", "2. Methane", "3. Nitrous Oxide", "4. All of the above"},
		{"1. Deforestation", "2. Industrial emissions", "3. Solar radiation", "4. All of the above"},
		{"1. Human activities", "2. Natural climate variability", "3. Combination of both", "4. None of the above"},
		{"1. Sea level rise", "2. Increased frequency of extreme weather events", "3. Biodiversity loss", "4. All of the above"},
		{"1. Renewable energy adoption", "2. Afforestation", "3. Carbon capture and storage", "4. All of the above"},
	}

	correctAnswers := []int{1, 4, 3, 4, 4}

	score := 0
	var userAnswer int

	for i := 0; i < 5; i++ {
		fmt.Println(questions[i])
		for _, option := range options[i] {
			fmt.Println(option)
		}
		fmt.Print("Enter your choice (1-4): ")
		fmt.Scanln(&userAnswer)
		if userAnswer == correctAnswers[i] {
			score++
		}
	}

	fmt.Println("Quiz complete. Your score is:", score)
	switch score {
	case 5:
		fmt.Println("Excellent")
	case 4:
		fmt.Println("Very good")
	default:
		fmt.Println("Time to brush up on your knowledge of global warming")
	}

	fmt.Println("For further reading, you can visit: [List of websites]")
}