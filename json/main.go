package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

func toJson(car Car) string {
	content, _ := json.Marshal(car)
	return string(content)
}

func toObject(car string) Car {
	var carReference Car
	toByte := []byte(car)
	json.Unmarshal(toByte, &carReference)
	return carReference
}

func main() {
	car := Car{
		Name: "Prisma",
		Year: 2020,
	}

	transform := toJson(car)
	revoke := toObject(transform)

	fmt.Println(transform)
	fmt.Println(revoke)
}
