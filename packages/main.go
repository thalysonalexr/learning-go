package main

import (
	"fmt"

	"github.com/thalysonalexr/learning-go/packages/pkg"
)

func main() {
	greeting, _ := pkg.Greeting("Hello")
	fmt.Println(greeting)
}
