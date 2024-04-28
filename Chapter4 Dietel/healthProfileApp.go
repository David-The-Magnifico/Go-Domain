package main

import (
	"fmt"
	"strconv"
)

type HealthProfile struct {
	firstName      string
	lastName       string
	gender         string
	birthYear      int
	heightInInches int
	weightInPounds int
}

func NewHealthProfile(firstName string, lastName string, gender string, birthYear int, heightInInches int, weightInPounds int) *HealthProfile {
	return &HealthProfile{
		firstName:      firstName,
		lastName:       lastName,
		gender:         gender,
		birthYear:      birthYear,
		heightInInches: heightInInches,
		weightInPounds: weightInPounds,
	}
}

func (hp *HealthProfile) CalculateAge(currentYear int) int {
	return currentYear - hp.birthYear
}

func (hp *HealthProfile) CalculateBMI() float64 {
	return (float64(hp.weightInPounds) / (float64(hp.heightInInches) * float64(hp.heightInInches))) * 703
}

func (hp *HealthProfile) CalculateMaxHeartRate(currentYear int) int {
	return 220 - hp.CalculateAge(currentYear)
}

func (hp *HealthProfile) CalculateTargetHeartRateRange(currentYear int) string {
	maxHeartRate := hp.CalculateMaxHeartRate(currentYear)
	minTargetHeartRate := int(float64(maxHeartRate) * 0.5)
	maxTargetHeartRate := int(float64(maxHeartRate) * 0.85)
	return strconv.Itoa(minTargetHeartRate) + " - " + strconv.Itoa(maxTargetHeartRate)
}

func main() {
	var firstName, lastName, gender string
	var birthYear, heightInInches, weightInPounds, currentYear int

	fmt.Print("Enter first name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Enter last name: ")
	fmt.Scanln(&lastName)

	fmt.Print("Enter gender: ")
	fmt.Scanln(&gender)

	fmt.Print("Enter year of birth: ")
	fmt.Scanln(&birthYear)

	fmt.Print("Enter height in inches: ")
	fmt.Scanln(&heightInInches)

	fmt.Print("Enter weight in pounds: ")
	fmt.Scanln(&weightInPounds)

	currentYear = 2023

	person := NewHealthProfile(firstName, lastName, gender, birthYear, heightInInches, weightInPounds)

	fmt.Println("Name: " + person.firstName + " " + person.lastName)
	fmt.Println("Age: " + strconv.Itoa(person.CalculateAge(currentYear)) + " years")
	fmt.Println("BMI: " + fmt.Sprintf("%.2f", person.CalculateBMI()))
	fmt.Println("Max Heart Rate: " + strconv.Itoa(person.CalculateMaxHeartRate(currentYear)) + " bpm")
	fmt.Println("Target Heart Rate Range: " + person.CalculateTargetHeartRateRange(currentYear) + " bpm")
}