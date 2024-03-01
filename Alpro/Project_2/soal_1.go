package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func permutation(x, y int) int {
	if x < y {
		return factorial(y) / factorial(y-x)
	}
	return factorial(x) / factorial(x-y)
}

func main() {
	var x, y int

	fmt.Scan(&x, &y)

	factorialX := factorial(x)
	factorialY := factorial(y)

	permutationResult := permutation(x, y)

	fmt.Println(factorialX, factorialY, permutationResult)
}
