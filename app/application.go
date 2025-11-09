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
	for _, v := range a {
		for _, k := range v {
			fmt.Println(k)
		}
	}
}
