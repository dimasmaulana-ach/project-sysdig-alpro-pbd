package main

import "fmt"

func main() {
	var dailyVisitors, prevVisitors, consecutiveIncreases, totalVisitors int
	consecutiveIncreases = 0

	fmt.Println("Museum Visitor Tracker")
	fmt.Println("Enter the number of daily visitors (enter -1 to stop):")

	for {
		fmt.Print("Hari ", consecutiveIncreases+1, ": ")
		fmt.Scanln(&dailyVisitors)

		if dailyVisitors == -1 {
			break
		}

		totalVisitors += dailyVisitors

		if dailyVisitors > prevVisitors {
			consecutiveIncreases++
			if consecutiveIncreases == 4 {
				fmt.Println("Program berhenti karena 3 kali kenaikan berturut-turut, yaitu")
				break
			}
		} else {
			consecutiveIncreases = 0
		}

		prevVisitors = dailyVisitors
	}

	fmt.Println("Jumlah pengunjung:", totalVisitors)
	fmt.Println("Program selesai.")
}
