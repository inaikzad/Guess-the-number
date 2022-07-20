package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GuessTheNumber() {

	fmt.Println("Write the highest number I'm allowed to think of: ")

	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')

	n, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	MyGuess := rand.Intn(n)

	fmt.Printf("Okay I thought of a number from 0 to %d ! Your guess is:\n", n)

	wrongGuessesCount := 0
	for {
		line, err = r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		guess, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			panic(err)
		}

		if guess == MyGuess {
			fmt.Println("You guessed the number correctly!")
			break

		} else {
			wrongGuessesCount++
			fmt.Println("Wrong guess! Try again...")
		}
	}

	fmt.Printf("You had %d wrong guesses.\n", wrongGuessesCount)
	if wrongGuessesCount < n/4 {
		fmt.Println("You got quite lucky!")
	} else if wrongGuessesCount < n/2 {
		fmt.Println("You did ok...")
	} else {
		fmt.Println("It took you a while...")
	}
}

func main() {

	var play string

	fmt.Println("Do you want to play a game of guess the number? (yes/no) ")
	fmt.Scanln(&play)

	if play == "yes" {
	StartGame:
		GuessTheNumber()
		var retry string

		fmt.Println("Do you want to play again? (yes/no)")
		fmt.Scanln(&retry)
		if retry == "yes" {
			goto StartGame
		} else {
			fmt.Println("Thanks for playing!")
		}
	} else {
		fmt.Println("Oh okay...")
	}

}
