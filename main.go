package main

import (
	"fmt"       // allow to print, read user input
	"math/rand" // random number generator (RNG)
	"strconv"
	"time" // allow 'time'
)

func main() {
	// new random number generator  - ensures that the random number sequence is different each time the program runs (RNG is independent of the global rand package.)
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	// get guess limit input as string
	var inputGuessLimit string
	fmt.Print("Enter a number as guess limit: ")
	fmt.Scanln(&inputGuessLimit)

	// validate guess limit whether it is int
	parsedGuessLimit, err := strconv.Atoi(inputGuessLimit)
	if err != nil {
		fmt.Print("Invalid input")
		return // TODO - temporary
	}

	// generate random number
	randNum := randGen.Intn(parsedGuessLimit) + 1

	// get user guess input as string
	var inputGuessNumber string
	fmt.Print("Your guess: ")
	fmt.Scanln(&inputGuessNumber)

	// validate user guess input whether it is int
	parsedGuessNumber, err := strconv.Atoi(inputGuessNumber)
	if err != nil {
		fmt.Printf("Invalid input %s", err)
		return // TODO - temporarily
	}

	fmt.Printf("Your input: %d - Limit: %d - Random number %d\n", parsedGuessNumber, parsedGuessLimit, randNum)

}

// Game Logic (simplified)
// initialize the game
// define range for guessing number (from .. to), this is the number to guess
// start guessing loop
// ask user input - guessed number
// if too high - say to guess lower
// if too low - say to guess higher
// if guessed number == randomly generated number - congratulate and finish game
