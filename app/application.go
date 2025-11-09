package app

import (
	"fmt"
)

func Run() {

	var message string
	var banner string

	message, banner = CheckArgs()
	fmt.Println(message, banner)
	a := LoadBanner(banner)
	for i := 0; i < 8; i++ {
		for _, v := range message {

			fmt.Print(a[v][i])
		}
		fmt.Print("\n")
	}
}
