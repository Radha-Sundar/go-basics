// Package comment provides a creative way to generate personalized fortunes.
//
// The fortunes are generated based on a fictional algorithm that combines the user's
// name and birthdate to create unique and positive messages. The fortunes are
// represented as strings.
package comment

import (
	"fmt"
	"math/rand"
	"time"
)

// Fortune represents a personalized fortune message.
type Fortune struct {
	Message string
}

// GenerateFortune generates a personalized fortune based on the user's name and birthdate.
//
// Parameters:
//
//	name: The name of the user.
//	birthdate: The birthdate of the user in the format "YYYY-MM-DD".
//
// Returns:
//
//	A Fortune struct containing the personalized fortune message.
//
// Example:
//
//	fortune := GenerateFortune("Alice", "1990-05-15")
//	fmt.Println("Your Fortune:", fortune.Message)
func GenerateFortune(name, birthdate string) Fortune {
	// A simple example of fortune generation
	rand.Seed(time.Now().UnixNano())
	fortuneMessages := []string{
		"Expect great things in your future,",
		"The stars align to bring you joy,",
		"Embrace the challenges ahead,",
		"Good fortune follows you,",
		"Your kindness will be rewarded,",
	}

	randomIndex := rand.Intn(len(fortuneMessages))

	message := fmt.Sprintf("%s %s! May your days be filled with happiness.", fortuneMessages[randomIndex], name)

	return Fortune{
		Message: message,
	}
}

// GetLuckyNumber generates a lucky number based on the user's birthdate.
//
// Parameters:
//
//	birthdate: The birthdate of the user in the format "YYYY-MM-DD".
//
// Returns:
//
//	An integer representing the lucky number.
//
// Example:
//
//	luckyNumber := GetLuckyNumber("1985-10-20")
//	fmt.Println("Your Lucky Number:", luckyNumber)
func GetLuckyNumber(birthdate string) int {
	// A fictional algorithm for generating a lucky number
	year, _, _ := parseBirthdate(birthdate)

	rand.Seed(time.Now().UnixNano())
	luckyNumber := rand.Intn(year) + 1

	return luckyNumber
}

// parseBirthdate extracts the year, month, and day components from a birthdate string.
func parseBirthdate(birthdate string) (int, int, int) {
	parsedDate, _ := time.Parse("2006-01-02", birthdate)
	return parsedDate.Year(), int(parsedDate.Month()), parsedDate.Day()
}
