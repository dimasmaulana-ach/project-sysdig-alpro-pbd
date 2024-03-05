package main

import (
	"fmt"
)

func isTriangle(a, b, c int) bool {
	return (a+b > c) && (a+c > b) && (b+c > a)
}

func main() {
	var side1, side2, side3 int
	fmt.Scan(&side1, &side2, &side3)

	if isTriangle(side1, side2, side3) {
		fmt.Println("Segitiga")
	} else {
		fmt.Println("Bukan segitiga")
	}
}
