package numberGuessing

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
)

func NumberGuessingGame(number int) (string, error) {

	min := 1
	max := 10
	randomN := rand.Intn(max-min+1) + min

	log.SetPrefix("GuessingGame: ")
	log.SetFlags(0)

	if number > 10 || number < 0 {
		err := errors.New("Number must be between 1 and 10, but received: " + strconv.Itoa(number))
		return "", err
	}

	log.Printf("Generated random number: %d", randomN)
	log.Printf("User's guess: %d", number)

	if randomN != number {
		return "You are a loser you lost", nil
	}

	return "You are amazing", nil
}
