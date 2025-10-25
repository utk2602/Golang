package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetValidatedInput(prompt string, validInputs []string) string {
	for {
		input := GetUserInput(prompt)
		if isValidInput(input, validInputs) {
			return input
		}
		fmt.Println("Invalid input. Please try again.")
	}
}

func isValidInput(input string, validInputs []string) bool {
	for _, valid := range validInputs {
		if strings.EqualFold(input, valid) {
			return true
		}
	}
	return false
}