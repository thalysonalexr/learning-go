package main

import (
	"fmt"
)

func Greating(name string) string {
	say := "Hello " + name
	return say
}

func main() {
	fmt.Println(Greating("Thalyson"))
}
