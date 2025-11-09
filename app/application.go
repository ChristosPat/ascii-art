package app

import (
	"fmt"
)

func Run() {

	var message string
	var banner string

	message, banner = CheckArgs()
	fmt.Println(message, banner)
}
