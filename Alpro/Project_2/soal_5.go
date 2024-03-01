package main

import "fmt"

func calculateDiscount(totalExpend int, member bool) int {
	var diskon int

	if totalExpend > 100000 && member {
		diskon = totalExpend / 10
	} else if totalExpend > 100000 && !member {
		diskon = totalExpend / 20
	} else {
		diskon = 0
	}

	return diskon
}

func main() {
	var totalExpend int
	var member bool

	fmt.Scan(&totalExpend, &member)

	discount := calculateDiscount(totalExpend, member)
	totalBelanjaAkhir := totalExpend - discount

	fmt.Println(totalBelanjaAkhir)
}
