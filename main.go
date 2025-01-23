package main

import (
	"fmt"       // allow to print, read user input
	"math/rand" // random number generator (RNG)
	"strconv"   // string to int
	// allow 'time'
)

func main() {
	// prompt and validate user guess limit
	var numberToGuess int = promptGuessLimit()

	// prompt and validate user guess
	var userGuess int = promptUserGuess(numberToGuess, "0")

	// check user guess against number to guess
	isGuessed(numberToGuess, userGuess)
}

// func to prompt and validate user guess limit
func promptGuessLimit() int {
	return rand.Intn(100) + 1 // random number between 1 to 100
}

// func to prompt and validate user guess
func promptUserGuess(guessLimit int, inputGuessNumber string) int {
	//fmt.Printf("\n[ check limit coming into promptUserGuess %d ]\n", guessLimit)
	//get user guess input as string
	fmt.Printf("\n> Please enter your guess between 1 to %d: ", guessLimit)
	fmt.Scanln(&inputGuessNumber)
	// validate user guess input whether it is int
	parsedGuessNumber, err := strconv.Atoi(inputGuessNumber)

	for err != nil {
		fmt.Print("\n ! Invalid input \n")
		fmt.Printf("\n> Please only enter a number between 1 to %d: ", guessLimit)
		fmt.Scanln(&inputGuessNumber)
		parsedGuessNumber, err = strconv.Atoi(inputGuessNumber)
	}
	return parsedGuessNumber
}

// func to check number to guess and user guess
func isGuessed(toGuess int, uNum int) {
	for toGuess != uNum { // equivalent of 'while' loop
		if uNum > toGuess {
			fmt.Print("\n :Try lower ⬇️\n")
		}
		if uNum < toGuess {
			fmt.Print("\n :Try higher ⬆️\n")
		}
		uNum = promptUserGuess(toGuess, strconv.Itoa(uNum)) // strconv.Itoa(intHere) - change int to str
	}
	fmt.Print("\n : You got it 🎉\n\n")
}

// Game Logic (simplified)
// initialize the game
// define range for guessing number (from .. to), this is the number to guess
// start guessing loop
// ask user input - guessed number
// if too high - say to guess lower
// if too low - say to guess higher
// if guessed number == randomly generated number - congratulate and finish game

// Terminal command
// go run main.go

// Git command
// # 1. ONLY IF there's changes made from github repo
// git pull origin main

// # 2. IF message requested, write something then
// 'Esc' for insert then ':wq' for write & quit

// # 3. Stage all changes
// git add .

// # 4. Commit the changes with a message
// git commit -m "your msg here"

// # 5. Push to the remote repository (to the 'main' branch)
// git push origin main
