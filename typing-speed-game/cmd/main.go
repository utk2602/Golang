package main

import (
	"fmt"
	"os"

	"github.com/utk2602/typing-speed-game/internal/game"
	"github.com/utk2602/typing-speed-game/internal/ui"
)

func main() {
	ui.DisplayWelcomeMessage()

	// Initialize the game
	g := game.NewGame()

	// Start the main game loop
	for {
		ui.DisplayMenu()
		choice := ui.GetUserInput("Enter your choice (1-3): ")

		switch choice {
		case "1":
			g.StartNewSession()
		case "2":
			g.DisplayAnalytics()
		case "3":
			ui.DisplayGoodbyeMessage()
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
