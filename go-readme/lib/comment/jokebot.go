// Package comment provides a creative way to generate jokes.
//
// The jokes are generated using a fictional algorithm that combines humorous
// elements and wordplay to create entertaining jokes. The resulting jokes are
// represented as strings.
package comment

import (
	"math/rand"
	"time"
)

// Joke represents a humorous joke.
type Joke struct {
	Text string
}

// GenerateRandomJoke generates a random joke with a touch of humor.
//
// Returns:
//
//	A Joke struct containing the randomly generated joke.
//
// Example:
//
//	randomJoke := GenerateRandomJoke()
//	fmt.Println("Random Joke:", randomJoke.Text)
func GenerateRandomJoke() Joke {
	// A simple example of joke generation
	rand.Seed(time.Now().UnixNano())
	jokeTypes := []string{"Knock-Knock", "One-Liner", "Pun"}

	randomJokeType := jokeTypes[rand.Intn(len(jokeTypes))]
	joke := ""

	switch randomJokeType {
	case "Knock-Knock":
		joke = generateKnockKnockJoke()
	case "One-Liner":
		joke = generateOneLinerJoke()
	case "Pun":
		joke = generatePunJoke()
	}

	return Joke{
		Text: joke,
	}
}

// generateKnockKnockJoke generates a knock-knock joke.
func generateKnockKnockJoke() string {
	// A fictional knock-knock joke
	return "Knock, knock.\nWho’s there?\nLettuce.\nLettuce who?\nLettuce in, it's freezing out here!"
}

// generateOneLinerJoke generates a one-liner joke.
func generateOneLinerJoke() string {
	// A fictional one-liner joke
	return "I told my wife she was drawing her eyebrows too high. She looked surprised."
}

// generatePunJoke generates a pun joke.
func generatePunJoke() string {
	// A fictional pun joke
	return "I used to be a baker because I kneaded dough."
}
