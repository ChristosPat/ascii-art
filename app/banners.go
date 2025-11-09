package app

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(file string) map[rune][]string {
	//	var banner map[rune][]string
	data, error := os.ReadFile("banners/" + file + ".txt")
	if error != nil {
		fmt.Println("Error reading file", error)
	}
	lines := strings.Split(string(data), "\n")
	banner := make(map[rune][]string)
	start := ' '
	for i := 1; i < len(lines); i += 9 {
		char := start
		start++
		banner[char] = lines[i : i+8]

	}

	return banner
}
