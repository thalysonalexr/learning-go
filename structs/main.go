package main

import "fmt"

type Car struct {
	name string
	year int
	km   float64
}

func (c *Car) drive() float64 {
	c.km += 5.67
	return c.km
}

func main() {
	var car Car = Car{
		name: "Fusion",
		year: 2011,
		km:   2800,
	}

	car.drive()
	fmt.Printf("%s andou %fkm\n", car.name, car.km)
	fmt.Println(car)
}
