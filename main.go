package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)
	myValue, err := strconv.ParseInt("7", 0, 64)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// map, los maps son estructuras de clave valor, se le especifica el tipo de dato de la clave normalmente string y el tipo de dato del valor
	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])

	// slice
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

}
