package main

import "fmt"

func calculateRegularCarFee(hours, minutes int) int {
	totalMinutes := hours*60 + minutes
	if totalMinutes <= 10 {
		return 0
	}
	extraHours := (totalMinutes - 10 + 59) / 60
	fee := 5000 + (extraHours-1)*6000
	return fee
}

func calculateTruckFee(hours, minutes int) int {
	return 2 * calculateRegularCarFee(hours, minutes)
}

func main() {
	var carType string
	var hours, minutes int

	fmt.Print("Enter car type (regular/truck): ")
	fmt.Scanln(&carType)
	fmt.Print("Enter parking duration (hours minutes): ")
	fmt.Scanln(&hours, &minutes)

	var fee int
	switch carType {
	case "regular":
		fee = calculateRegularCarFee(hours, minutes)
	case "truck":
		fee = calculateTruckFee(hours, minutes)
	default:
		fmt.Println("Invalid car type")
		return
	}

	fmt.Println("Parking fee:", fee, "rupiah")
}
