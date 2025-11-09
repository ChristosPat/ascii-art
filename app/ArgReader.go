package app

import (
	"fmt"
	"os"
	"slices"
)

// Check the arguments and gives the text and banner
func CheckArgs() (string, string) {
	var banner string
	if len(os.Args) < 2 {
		fmt.Println("Usage go run . <text> <banner>")
		os.Exit(1)
	}
	message := os.Args[1]
	if len(os.Args) == 2 {
		banner = "standard"
	} else {
		banner = os.Args[2]
		if !CheckBanner(banner) {
			fmt.Println("Allowed banners are: shadow - standard - thinkertoy .")
			os.Exit(2)
		}

	}
	return message, banner
}

// Check the allowed banners
func CheckBanner(banner string) bool {
	bannerName := []string{
		"shadow", "standard", "thinkertoy"}

	return slices.Contains(bannerName, banner)

}
