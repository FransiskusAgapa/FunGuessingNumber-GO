package main

import (
	"fmt"       // allow to print, read user input
	"math/rand" // random number generator (RNG)
	"strconv"   // string to int
	// allow 'time'
)

// TODO:
// # UX
// # 1. Provide hint: how 'close' their guess to the actual number
// # 2. Provide Guess History: a way to keep track of previous guess & display it
// # 3. Set timer/timeout: countdown/limit the number of guess
// # Interactive Feedback
// # 4. Progressive Feedback: Instead of just saying "Try lower" or "Try higher," include more engaging responses, like "You're getting warmer!" or "You're freezing cold!"
// # 5. Sound Effects: Use sound effects to provide feedback on user actions, such as a "ding" when they guess correctly or a "buzz" for incorrect guesses.
// # Personalization
// # 6. Set player's name
// # 7. Leaderboard: Implement a scoring system where players earn points based on how quickly they guess the number, and display a leaderboard of top scores.
// # Error Handling and Input Validation
// # 8. Provide friendly error messages for invalid inputs.
// # 9. Retry Option: After the game ends, give the player the option to play again without restarting the entire program.
// # Story / Theme
// # 10. Themed Game: Add a narrative or theme to the game, such as guessing a number to unlock a treasure chest or as part of a larger adventure. This can increase engagement by giving the game context and meaning.
// # 11. Progressive Challenges: Introduce new challenges or mini-games as the user progresses through levels or wins multiple rounds, such as more difficult numbers to guess or added obstacles.
// # Multiplayer option
// # 12. Competitive Mode: Allow multiple players to take turns guessing the number and see who can guess it in the fewest attempts.
// # 13. Tracking Statistic
// # 14. Guessing Stats: Track statistics such as the average number of guesses per game, the longest streak of correct guesses, and the time taken to guess the number. Display these stats at the end of the game to encourage improvement.

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
