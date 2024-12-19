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
	fmt.Print("\n> Please enter a number as guess limit: ")
	fmt.Scanln(&inputGuessLimit)

	// validate guess limit whether it is int
	parsedGuessLimit, err := strconv.Atoi(inputGuessLimit)
	for err != nil {
		fmt.Print("\n ! Invalid input\n")

		// keep asking for correct input
		fmt.Print("\n> Please only enter a number guess limit: ")
		fmt.Scanln(&inputGuessLimit)
		parsedGuessLimit, err = strconv.Atoi(inputGuessLimit)
	}

	// generate random number
	randNumToGuess := randGen.Intn(parsedGuessLimit) + 1

	// get user guess input as string
	var inputGuessNumber string
	fmt.Printf("\n> Please enter your guess between 1 to %d: ", parsedGuessLimit)
	fmt.Scanln(&inputGuessNumber)

	// validate user guess input whether it is int
	parsedGuessNumber, err := strconv.Atoi(inputGuessNumber)
	for err != nil {
		fmt.Print("\n ! Invalid input\n")

		fmt.Printf("\n> Please only enter your guess between 1 to %d: ", parsedGuessLimit)
		fmt.Scanln(&inputGuessNumber)
		parsedGuessNumber, err = strconv.Atoi(inputGuessNumber)
	}

	//var isGuessedCorrect bool = false
	//for !isGuessedCorrect {
	if parsedGuessNumber > randNumToGuess {
		fmt.Print("\n :Try lower ⬇️: ")
	}
	if parsedGuessNumber < randNumToGuess {
		fmt.Print("\n :Try higher ⬆️: ")
	}
	if parsedGuessNumber == randNumToGuess {
		fmt.Print("\n :You got it 🎉")
		//isGuessedCorrect = true
	}
	//}
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
// # 1. Stage all changes
// git add .

// # 2. Commit the changes with a message
// git commit -m "Fixed bug in the login form"

// # 3. Push to the remote repository (to the 'main' branch)
// git push origin main
