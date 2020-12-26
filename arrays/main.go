package main

import "fmt"

func main() {
	const MAX = 5

	// static arrays
	var values [MAX]int

	values[0] = 1
	values[1] = 2

	fmt.Println(values)

	// initialize static arrays
	values2 := [MAX]int{0, 1, 2, 3, 4}

	fmt.Println(values2)

	// array list
	values3 := make([]int, 0)

	copyValues := append(values3, 1, 2, 3)
	fmt.Println(copyValues)

	// initialize slices
	oneSlice := []int{1, 2}
	fmt.Println(oneSlice)

	// maps
	type Coordenates struct {
		Lat, Long float64
	}

	m := make(map[string]Coordenates)

	m["Bell Labs"] = Coordenates{
		40.68433, -74.39967,
	}
	m["My House"] = Coordenates{
		-16.4796311, -54.6426169,
	}

	fmt.Println(m)
}
