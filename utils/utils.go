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
func DisplayWarningMessage(keyType, privateKey, ksPath string) {
	// Clear the screen
	cmd := exec.Command("clear") // For Windows use "cls"
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("failed to clear the screen: ", err)
	}

	// Display warning message with ASCII art and colored text
	fmt.Println("\033[1m\x1b[36m#########################################################################\033[0m")
	fmt.Println("\033[1m\x1b[36m#  _        _____    _____   _____    _____   _     _   _____   ______  #\033[0m")
	fmt.Println("\033[1m\x1b[36m# (_)      (_____)  (_____) (_____)  (_____) (_)   (_) (_____) (______) #\033[0m")
	fmt.Println("\033[1m\x1b[36m# (_)     (_)___(_)(_)  ___ (_)__(_)(_)___(_)(__)_ (_)(_)  ___ (_)__    #\033[0m")
	fmt.Println("\033[1m\x1b[36m# (_)     (_______)(_) (___)(_____) (_______)(_)(_)(_)(_) (___)(____)   #\033[0m")
	fmt.Println("\033[1m\x1b[36m# (_)____ (_)   (_)(_)___(_)( ) ( ) (_)   (_)(_)  (__)(_)___(_)(_)____  #\033[0m")
	fmt.Println("\033[1m\x1b[36m# (______)(_)   (_) (_____) (_)  (_)(_)   (_)(_)   (_) (_____) (______) #\033[0m")
	fmt.Println("\033[1m\x1b[36m#########################################################################\033[0m")

	fmt.Println("")
	fmt.Println(strings.ToUpper(keyType) + " Private Key Generated Successfully 🎉")
	border := strings.Repeat("=", len(privateKey)+6)
	fmt.Println("\033[1m\x1b[31m" + border + "\033[0m")
	fmt.Println("\x1b[36m|  " + privateKey + "  |\033[0m")
	fmt.Println("\033[1m\x1b[31m" + border + "\033[0m")
	fmt.Println("")
	fmt.Println("\033[1m\x1b[33m🔑  WARNING: Make sure to copy this private key securely and never share it with anyone!\033[0m")
	fmt.Println("🔑  Keystore file saved at: " + ksPath)
}
