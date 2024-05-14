package utils

import (
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"strings"
)

// ConvertBLSKey converts a BLS public key from a string to a big.Int array.
func ConvertBLSKey(blsPubKey []byte) ([2]*big.Int, error) {
	// Parse BLS public key
	if len(blsPubKey) != 64 {
		return [2]*big.Int{}, fmt.Errorf("invalid BLS public key")
	}

	var pubKey [2]*big.Int
	pubKey[0] = new(big.Int).SetBytes(blsPubKey[:32])
	pubKey[1] = new(big.Int).SetBytes(blsPubKey[32:])

	return pubKey, nil
}

// DisplayWarningMessage displays a warning message with ASCII art.
func DisplayWarningMessage(keyType, privateKey, ksPath string) error {
	// Clear the screen
	cmd := exec.Command("less", "-R") // For Windows use "cls"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdin pipe: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}

	msg := ""
	// Display warning message with ASCII art and colored text
	msg += fmt.Sprintln("\033[1m\x1b[36m#########################################################################\033[0m")
	msg += fmt.Sprintln("\033[1m\x1b[36m#  _        _____    _____   _____    _____   _     _   _____   ______  #\033[0m")
	msg += fmt.Sprintln("\033[1m\x1b[36m# (_)      (_____)  (_____) (_____)  (_____) (_)   (_) (_____) (______) #\033[0m")
	msg += fmt.Sprintln("\033[1m\x1b[36m# (_)     (_)___(_)(_)  ___ (_)__(_)(_)___(_)(__)_ (_)(_)  ___ (_)__    #\033[0m")
	msg += fmt.Sprintln("\033[1m\x1b[36m# (_)     (_______)(_) (___)(_____) (_______)(_)(_)(_)(_) (___)(____)   #\033[0m")
	msg += fmt.Sprintln("\033[1m\x1b[36m# (_)____ (_)   (_)(_)___(_)( ) ( ) (_)   (_)(_)  (__)(_)___(_)(_)____  #\033[0m")
	msg += fmt.Sprintln("\033[1m\x1b[36m# (______)(_)   (_) (_____) (_)  (_)(_)   (_)(_)   (_) (_____) (______) #\033[0m")
	msg += fmt.Sprintln("\033[1m\x1b[36m#########################################################################\033[0m")

	msg += fmt.Sprintln("")
	msg += fmt.Sprintln(strings.ToUpper(keyType) + " Private Key Generated Successfully 🎉")
	border := strings.Repeat("=", len(privateKey)+6)
	msg += fmt.Sprintln("\033[1m\x1b[31m" + border + "\033[0m")
	msg += fmt.Sprintln("\x1b[36m|  " + privateKey + "  |\033[0m")
	msg += fmt.Sprintln("\033[1m\x1b[31m" + border + "\033[0m")
	msg += fmt.Sprintln("")
	msg += fmt.Sprintln("\033[1m\x1b[33m🔑  WARNING: Make sure to copy this private key securely and never share it with anyone!\033[0m")
	msg += fmt.Sprintln("🔑  Keystore file saved at: " + ksPath)

	// Write the message to the command
	if _, err = stdin.Write([]byte(msg)); err != nil {
		return fmt.Errorf("failed to write to stdin: %w", err)
	}

	// Close the stdin pipe
	if err := stdin.Close(); err != nil {
		return fmt.Errorf("failed to close stdin: %w", err)
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to wait for command: %w", err)
	}

	return nil
}
