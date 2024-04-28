package main

import (
	"fmt"
)

type HeartRates struct {
	firstName  string
	lastName   string
	birthYear  int
}

func NewHeartRates(firstName string, lastName string, birthYear int) *HeartRates {
	return &HeartRates{firstName, lastName, birthYear}
}

func (hr *HeartRates) calculateAge(currentYear int) int {
	return currentYear - hr.birthYear
}

func (hr *HeartRates) calculateMaxHeartRate(currentYear int) int {
	return 220 - hr.calculateAge(currentYear)
}

func (hr *HeartRates) calculateTargetHeartRateRange(currentYear int) string {
	maxHeartRate := hr.calculateMaxHeartRate(currentYear)
	minTargetHeartRate := int(float64(maxHeartRate) * 0.5)
	maxTargetHeartRate := int(float64(maxHeartRate) * 0.85)
	return fmt.Sprintf("%d - %d", minTargetHeartRate, maxTargetHeartRate)
}

func main() {
	var firstName, lastName string
	var birthYear, currentYear int

	fmt.Print("Enter first name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Enter last name: ")
	fmt.Scanln(&lastName)

	fmt.Print("Enter year of birth: ")
	fmt.Scanln(&birthYear)

	currentYear = 2023

	person := NewHeartRates(firstName, lastName, birthYear)

	fmt.Printf("Name: %s %s\n", person.firstName, person.lastName)
	fmt.Printf("Age: %d years\n", person.calculateAge(currentYear))
	fmt.Printf("Max Heart Rate: %d bpm\n", person.calculateMaxHeartRate(currentYear))
	fmt.Printf("Target Heart Rate Range: %s bpm\n", person.calculateTargetHeartRateRange(currentYear))
}