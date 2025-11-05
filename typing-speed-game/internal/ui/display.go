package ui

import (
	"fmt"
	"strings"
)

func DisplayWelcomeMessage() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("  Welcome to the Typing Speed Game!")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Test your typing speed and improve your skills.")
	fmt.Println("Type the given text as quickly and accurately as possible.")
	fmt.Println(strings.Repeat("=", 50) + "\n")
}

func DisplayMenu() {
	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("MENU:")
	fmt.Println("  1. Start New Game")
	fmt.Println("  2. View Analytics")
	fmt.Println("  3. Exit")
	fmt.Println(strings.Repeat("-", 50))
}

func DisplayResults(speed float64, accuracy float64) {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("  RESULTS")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("  Typing Speed: %.2f WPM\n", speed)
	fmt.Printf("  Accuracy: %.2f%%\n", accuracy)
	fmt.Println(strings.Repeat("=", 50))
}

func DisplayAnalytics(averageSpeed float64, bestSpeed float64, worstSpeed float64) {
	fmt.Println("  TYPING SPEED ANALYTICS")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("  Average Speed: %.2f WPM\n", averageSpeed)
	fmt.Printf("  Best Speed: %.2f WPM\n", bestSpeed)
	fmt.Printf("  Worst Speed: %.2f WPM\n", worstSpeed)
}

func DisplayGoodbyeMessage() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("  Thank you for playing the Typing Speed Game!")
	fmt.Println("  Goodbye!")
	fmt.Println(strings.Repeat("=", 50) + "\n")
}
