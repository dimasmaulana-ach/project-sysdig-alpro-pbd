package main

import (
	"fmt"
)

func main() {
	const (
		targetScore = 70
		numSets     = 10
	)

	var totalScore, setsCompleted int
	var scores [3]int

	fmt.Println("Mencatat Kekonsistenan Memanah")
	fmt.Println("Skor per lingkaran dari terluar hingga terdalam: 1, 4, 7, 10")
	fmt.Println("Masukkan skor per set latihan (3 kali kesempatan memanah per set), pisahkan dengan spasi:")
	for setsCompleted = 1; setsCompleted <= numSets; setsCompleted++ {
		fmt.Printf("Set ke-%d: ", setsCompleted)
		fmt.Scanln(&scores[0], &scores[1], &scores[2])

		setScore := scores[0] + scores[1] + scores[2]
		totalScore += setScore

		if setScore == 30 || totalScore >= targetScore {
			break
		}
	}

	fmt.Println("SELESAI")
	if totalScore >= targetScore {
		fmt.Printf("Tercapai pada set ke-%d\n", setsCompleted)
	} else {
		fmt.Println("Atlet tidak mencapai target dalam 10 set.")
	}
}
