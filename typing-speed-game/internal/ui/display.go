package ui

import (
	"fmt"
)

func DisplayWelcomeMessage() {
	fmt.Println("Welcome to the Typing Speed Game!")
	fmt.Println("Test your typing speed and improve your skills.")
	fmt.Println("Type the given text as quickly and accurately as possible.")
	fmt.Println("Press 'Enter' to start the game.")
}

func DisplayResults(speed float64, accuracy float64) {
	fmt.Printf("Your typing speed: %.2f WPM\n", speed)
	fmt.Printf("Your accuracy: %.2f%%\n", accuracy)
}

func DisplayAnalytics(averageSpeed float64, bestSpeed float64, worstSpeed float64) {
	fmt.Println("Typing Speed Analytics:")
	fmt.Printf("Average Speed: %.2f WPM\n", averageSpeed)
	fmt.Printf("Best Speed: %.2f WPM\n", bestSpeed)
	fmt.Printf("Worst Speed: %.2f WPM\n", worstSpeed)
}

func DisplayGoodbyeMessage() {
	fmt.Println("Thank you for playing the Typing Speed Game!")
	fmt.Println("Goodbye!")
}