package main

import (
	"fmt"
	"os"
	"typing-speed-game/internal/game"
	"typing-speed-game/internal/ui"
)

func main() {
	fmt.Println("Welcome to the Typing Speed Game!")
	
	// Initialize the game
	g := game.NewGame()

	// Start the main game loop
	for {
		ui.DisplayMenu()
		choice := ui.GetUserInput()

		switch choice {
		case "1":
			g.StartNewSession()
		case "2":
			g.DisplayAnalytics()
		case "3":
			fmt.Println("Exiting the game. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}