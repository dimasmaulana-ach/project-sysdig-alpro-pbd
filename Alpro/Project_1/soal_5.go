package main

import "fmt"

func main() {
	var N int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scanln(&N)

	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
