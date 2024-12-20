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
	var numberToGuess int = promptGuessLimit(randGen)
	var userGuess int = promptUserGuess(numberToGuess)
	//fmt.Printf("\n-ToGuess: %d\n-YourInput: %d", numberToGuess, userGuess)
	isGuessed(numberToGuess, userGuess)
}

// func to prompt and validate user guess limit
func promptGuessLimit(rng *rand.Rand) int {
	var randNumToGuess int
	var inputGuessLimit string
	fmt.Print("\n> Please enter a number as guess limit: ")
	fmt.Scanln(&inputGuessLimit)
	parsedGuessLimit, err := strconv.Atoi(inputGuessLimit)

	// keep asking for correct input (an int)
	for err != nil {
		fmt.Print("\n ! Invalid input\n")
		fmt.Print("\n> Please only enter a number as guess limit: ")
		fmt.Scanln(&inputGuessLimit)
		parsedGuessLimit, err = strconv.Atoi(inputGuessLimit)
	}
	randNumToGuess = rng.Intn(parsedGuessLimit) + 1

	//fmt.Printf("\n[ check guess limit Before leaving promptGuessLimit: %d ]\n", randNumToGuess)

	return randNumToGuess
}

// func to prompt and validate user guess
func promptUserGuess(guessLimit int) int {
	//fmt.Printf("\n[ check limit coming into promptUserGuess %d ]\n", guessLimit)

	//get user guess input as string
	var inputGuessNumber string
	fmt.Printf("\n> Please enter your guess between 1 to %d: ", guessLimit)
	fmt.Scanln(&inputGuessNumber)
	// validate user guess input whether it is int
	parsedGuessNumber, err := strconv.Atoi(inputGuessNumber)

	for err != nil {
		fmt.Print("\n ! Invalid input \n")
		fmt.Printf("\n> Please only enter your guess between 1 to %d: ", guessLimit)
		fmt.Scanln(&inputGuessNumber)
		parsedGuessNumber, err = strconv.Atoi(inputGuessNumber)
	}
	return parsedGuessNumber
}

// func to check number to guess and user guess
func isGuessed(toGuess int, uNum int) bool {
	if uNum > toGuess {
		fmt.Print("\n :Try lower ⬇️\n")
	}
	if uNum < toGuess {
		fmt.Print("\n :Try higher ⬆️\n")
	}
	if uNum == toGuess {
		fmt.Print("\n :You got it 🎉\n")
		return true
	}
	return false
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
// git commit -m "Fixed bug in the login form"

// # 5. Push to the remote repository (to the 'main' branch)
// git push origin main
