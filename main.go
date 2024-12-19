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
	fmt.Print("\nPlease enter a number as guess limit: ")
	fmt.Scanln(&inputGuessLimit)

	// validate guess limit whether it is int
	parsedGuessLimit, err := strconv.Atoi(inputGuessLimit)
	for err != nil {
		fmt.Print("]n> Invalid input\n")
		// keep asking for correct input
		fmt.Print("Please only enter a number guess limit: ")
		fmt.Scanln(&inputGuessLimit)
		parsedGuessLimit, err = strconv.Atoi(inputGuessLimit)
	}

	// generate random number
	randNum := randGen.Intn(parsedGuessLimit) + 1

	// get user guess input as string
	var inputGuessNumber string
	fmt.Printf("Please enter your guess between 1 to %d: ", parsedGuessLimit)
	fmt.Scanln(&inputGuessNumber)

	// validate user guess input whether it is int
	parsedGuessNumber, err := strconv.Atoi(inputGuessNumber)
	if err != nil {
		fmt.Printf("Invalid input %s", err)
		return // TODO - temporarily
	}

	fmt.Printf("Your Guess: %d\n-Limit: %d\n- Random number to guess %d\n", parsedGuessNumber, parsedGuessLimit, randNum)

}

// Game Logic (simplified)
// initialize the game
// define range for guessing number (from .. to), this is the number to guess
// start guessing loop
// ask user input - guessed number
// if too high - say to guess lower
// if too low - say to guess higher
// if guessed number == randomly generated number - congratulate and finish game

// Git command
// add all changes and commit => git commit -am "Your commit message"
// puss changes to github => git push
