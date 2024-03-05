package main

import "fmt"

func sumFibonacci(N int) int {
	if N <= 0 {
		return 0
	}

	fibSum := 0
	prev, curr := 0, 1

	for i := 1; i <= N; i++ {
		fibSum += prev
		prev, curr = curr, prev+curr
	}

	return fibSum
}

func main() {
	var N int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scanln(&N)

	result := sumFibonacci(N)
	fmt.Println(result)
}
