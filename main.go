package main

import (
	"/game/questions"
	"fmt"
	"os"
	"strconv"
)

// Function to get user input
func getUserInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

// Main Menu function
func main() {
	fmt.Println("Welcome to the History Quiz Game!")

	for {
		fmt.Println("1. Start Quiz")
		fmt.Println("2. View High Scores")
		fmt.Println("3. Exit")
		fmt.Print("> ")

		choice := getUserInput()
		switch choice {
		case "1":
			fmt.Println("Choose Epoch:")
			fmt.Println("1. Ancient History")
			fmt.Print("> ")

			epochChoice := getUserInput()
			var questionsList []questions.Question

			// Select questions based on epoch
			switch epochChoice {
			case "1":
				questionsList = questions.AncientHistoryQuestions
			default:
				fmt.Println("Invalid epoch choice.")
				continue
			}

			// Debug: Print the number of questions loaded
			fmt.Printf("Number of questions loaded: %d\n", len(questionsList))

			// Start quiz stage if enough questions are available
			if len(questionsList) >= 10 {
				startQuizStage(questionsList)
			} else {
				fmt.Println("Not enough questions available for this epoch.")
			}

		case "2":
			displayHighScores("Ancient History") // Implement high score display
		case "3":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, please select a valid option.")
		}
	}
}

// Function to start the quiz stage
func startQuizStage(questionsList []questions.Question) {
	correctAnswers := 0

	// Ask 10 random questions
	for i := 0; i < 10; i++ {
		question := questionsList[i]
		fmt.Println(question.Question)
		for j, choice := range question.Choices {
			fmt.Printf("%d. %s\n", j+1, choice)
		}
		fmt.Print("> ")

		answer := getUserInput()
		answerIndex, err := strconv.Atoi(answer)
		if err != nil || answerIndex < 1 || answerIndex > len(question.Choices) {
			fmt.Println("Invalid input. Please choose a valid option.")
			i-- // Retry the current question
			continue
		}

		// Check if the answer is correct
		if answerIndex-1 == question.Correct {
			correctAnswers++
		}
	}

	// Display results
	fmt.Printf("You got %d/10 correct answers!\n", correctAnswers)
	if correctAnswers >= 7 {
		fmt.Println("Congratulations! You passed.")
		startBattlePhase() // Implement this function for the battle phase
	} else {
		fmt.Println("Sorry, you didn't pass. Try again!")
		fmt.Println("1. Retry Quiz")
		fmt.Println("2. Exit")
		retryChoice := getUserInput()

		if retryChoice == "1" {
			startQuizStage(questionsList)
		} else {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}
}

// Function to display high scores (dummy implementation for now)
func displayHighScores(epoch string) {
	fmt.Printf("Displaying high scores for %s...\n", epoch)
	// Add actual high score logic here
}

// Function to start the battle phase (dummy implementation for now)
func startBattlePhase() {
	fmt.Println("Entering the battle phase...")
	// Add the battle phase logic here
}
