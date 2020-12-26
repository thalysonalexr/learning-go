package main

import (
	"fmt"

	"github.com/fatih/color"
	"rsc.io/quote"
)

func Hello() string {
	return "Hello"
}

func main() {
	color.Cyan(quote.Glass())
	fmt.Println(Hello())
}
