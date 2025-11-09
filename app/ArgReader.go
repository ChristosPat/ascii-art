package app

import (
	"fmt"
	"os"
)

// Check the arguments and gives the text and banner
func CheckArgs() (string, string) {
	var banner string
	if len(os.Args) < 2 {
		fmt.Println("Usage go run . <text> <banner>")
	}
	message := os.Args[1]
	if len(os.Args) < 2 {
		banner = "standard"
	} else {
		banner = os.Args[2]
	}
	return message, banner
}
