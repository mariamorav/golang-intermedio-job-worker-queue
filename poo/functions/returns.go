package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func sumVariadica(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}

	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// retornos con nombre
func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(1, 3))
	fmt.Println(sumVariadica(1, 4, 5))

	printNames("Maria", "Ana", "Andrea")

	fmt.Println(getValues(2))
}
