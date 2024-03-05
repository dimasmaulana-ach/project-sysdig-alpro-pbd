package main

import (
	"fmt"
)

func convertToSeconds(hour, minute, second int) int {
	totalSeconds := (hour * 3600) + (minute * 60) + second
	return totalSeconds
}

func main() {
	var hour, minute, second int
	fmt.Println("Masukkan waktu dalam format jam, menit, detik:")
	fmt.Scan(&hour, &minute, &second)

	totalSeconds := convertToSeconds(hour, minute, second)
	fmt.Println("Hasil konversi:", totalSeconds, "detik")
}
