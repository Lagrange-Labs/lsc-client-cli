package utils

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

// StringPrompt prompts the user for a string input.
func StringPrompt(label string) (string, error) {
	var input string
	for {
		fmt.Fprint(os.Stderr, label)
		if _, err := fmt.Fscanln(os.Stdin, &input); err != nil {
			return "", err
		}
		if len(input) > 0 {
			break
		}
		fmt.Fprintln(os.Stderr, "Input cannot be empty")
	}
	return input, nil
}

// PasswordPrompt prompts the user for a password input.
func PasswordPrompt(label string) (string, error) {
	var password string
	for {
		fmt.Fprint(os.Stderr, label)
		bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return "", err
		}
		password = string(bytePassword)
		if len(password) > 0 {
			break
		}
		fmt.Fprintln(os.Stderr, "Password cannot be empty")
	}
	fmt.Fprintln(os.Stderr)
	return password, nil
}

// ConfirmPrompt prompts the user for a confirmation input.
func ConfirmPrompt(label string) (bool, error) {
	var input string
	for {
		fmt.Fprint(os.Stderr, label)
		if _, err := fmt.Fscanln(os.Stdin, &input); err != nil {
			return false, err
		}
		switch strings.ToLower(strings.Trim(input, " ")) {
		case "y", "yes":
			return true, nil
		case "n", "no":
			return false, nil
		default:
			fmt.Fprintln(os.Stderr, "Invalid input. Please enter 'y' or 'n'")
		}
	}
}
