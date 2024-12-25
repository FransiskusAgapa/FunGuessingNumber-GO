# A Game in GO - Guess Number

The purpose of this assignment are:
1. Learn yet another programming language
2. Practice creativity
3. Apply the concepts of  Data Types

## Description of Guessing Number Game
Let's see how luck you are guessing a random number! Set the range of number for your to guess a (guaranteed) random number. 

## How to run
It's simple!
* Install GO, TBA

## How to play (screenshots of the gameplay)

   ### A (YouTube) video showing the game running and explaining how it goes.

## Data types & Concepts
Data Types:  
`int`: var numberToGuess int and var userGuess int are both integers, representing the random number to guess and the user's guess.

`string`: var `inputGuessLimit` string and var `inputGuessNumber` string store user inputs as strings before converting them to integers.

`error`: For example, err is used to check if the conversion from string to integer `(strconv.Atoi())` was successful or not.

`*rand.Rand`: A pointer to a Rand object from the math/rand package. This is used for generating random numbers. The `randGen` variable is an instance of `Rand`, created with a specific seed (`rand.New(rand.NewSource(time.Now().UnixNano()))`), ensuring that the random sequence is different each time the program runs.

`fmt`: for formatted I/O operations, such as printing output (`fmt.Print`, `fmt.Printf`) and reading user input (`fmt.Scanln`).

`strconv`: The `strconv` package provides functions for converting between string and numeric types. `strconv.Atoi()` is used to convert a string to an integer, and `strconv.Itoa()` converts an integer to a string.

`time`: Work with time-related functionality. In this code, `time.Now().UnixNano()` is used to generate a unique random seed based on the current time.

Concepts:
Function Definitions: Using `func`. For example, `func main()` is the entry point of the program, and other functions like `promptGuessLimit`, `promptUserGuess`, and `isGuessed` are defined to handle specific tasks.

Pointers: The `randGen` variable is a pointer to a `rand.Rand` instance. Pointers allow you to reference and manipulate the actual memory location of a variable, which is useful when you want to pass large data structures efficiently or modify data within functions.

Error Handling: The code uses `err` to handle errors that occur when converting a string input to an integer using `strconv.Atoi()`. If the conversion fails (e.g., if the user enters a non-numeric string), it will enter a loop prompting the user for valid input.

Loops: The for loop is used in multiple places - In `promptGuessLimit` and `promptUserGuess`, to keep asking the user for valid input until a valid number is provided.
In `isGuessed`, a for loop is used to repeatedly check the user's guess until they guess the correct number.

Conditional Statements: if statements are used to check whether the user's guess is higher or lower than the target number and provide feedback to guide the user.

Type Conversion: `strconv.Atoi()` to convert strings to integers and `strconv.Itoa()` to convert integers to strings. This is necessary because user input is always read as a string and needs to be converted for numeric comparisons.

Random Number Generation: The rand package is used to generate random numbers. The `rand.NewSource(time.Now().UnixNano())` is used to create a new source of randomness based on the current time, to make different random numbers each time the program runs.

## Difficulties and Solutions (challenges found & how it was overcome)
* Difficulties: 
** Figure out the data type to be used
** Figure out how user input will be taken and compared
  
* Solutions:
** Make a step-by-step outline of how the game plays out
  
## The Good, the Bad and the Ugly (What was loved about this experience, what was bad, and what did was disliked)
* Good: 
** It's a fun and intuitive game, using emoji to react to user guess is FUN

* Bad: 
** -

* Ugly:
** -
  
## Learning Experience
* Lessons: 
** Learned how different types are instantiated in GO.
** Learned how error is useful to validate the state of a value
** Learned how to use cool conversion int to string, vice versa.
** Learned to use time-related functionality to generate unique random seed based on the current time


