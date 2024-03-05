// soal 1

//package main
//
//import "fmt"
//
//func konversiKoin(uangLumin int) (int, int, int, int, int) {
//	quantumShard := uangLumin / (20 * 2 * 3 * 10)
//	sisa := uangLumin % (20 * 2 * 3 * 10)
//
//	galactum := sisa / (20 * 2 * 3)
//	sisa %= (20 * 2 * 3)
//
//	nebulaCrown := sisa / (20 * 2)
//	sisa %= (20 * 2)
//
//	stellaris := sisa / 20
//	sisa %= 20
//
//	return quantumShard, galactum, nebulaCrown, stellaris, sisa
//}
//
//func main() {
//	var uangLumin int
//	fmt.Scan(&uangLumin)
//	quantumShard, galactum, nebulaCrown, stellaris, sisa := konversiKoin(uangLumin)
//	fmt.Println(quantumShard, galactum, nebulaCrown, stellaris, sisa)
//}

//soal 2
//package main
//
//import "fmt"
//
//
//func cetakPola(n int) {
//	if n%2 == 0 {
//		fmt.Println("Bilangan harus ganjil")
//		return
//	}
//
//	for i := 0; i < n; i++ {
//		for j := 0; j < n; j++ {
//			if j == i {
//				fmt.Print(" ")
//			} else if j == n-1-i {
//				fmt.Print("*")
//			} else {
//				fmt.Print("*")
//			}
//		}
//		fmt.Println()
//	}
//}
//
//func main() {
//	var n int
//	fmt.Scan(&n)
//	cetakPola(n)
//}

//soal 3
//
//package main
//
//import "fmt"
//
//func tentukanKuadran(x, y float64) string {
//	if x > 0 && y > 0 {
//		return "Kuadran I"
//	} else if x < 0 && y > 0 {
//		return "Kuadran II"
//	} else if x < 0 && y < 0 {
//		return "Kuadran III"
//	} else if x > 0 && y < 0 {
//		return "Kuadran IV"
//	} else {
//		return "Titik berada pada sumbu"
//	}
//}
//
//func main() {
//	var x, y float64
//	fmt.Scan(&x, &y)
//	fmt.Println(tentukanKuadran(x, y))
//}

//soal 4

//package main
//
//import "fmt"
//
//func main() {
//	var n, kekuatan, kecepatan int
//	fmt.Scan(&n)
//	fmt.Scan(&kekuatan, &kecepatan)
//
//	terkalahkan := 0
//
//	for i := 0; i < n; i++ {
//		var kekuatanLawan, kecepatanLawan int
//		fmt.Scan(&kekuatanLawan, &kecepatanLawan)
//
//		if kekuatan >= 3 && kecepatan >= 3 {
//			if kekuatanLawan > kekuatan && kecepatanLawan > kecepatan {
//				continue
//			} else if kekuatanLawan >= kecepatanLawan && kekuatanLawan > kekuatan {
//				kecepatan -= 6
//				if kecepatan < 3 {
//					kecepatan = 3
//				}
//				terkalahkan++
//			} else if kecepatanLawan >= kekuatanLawan && kecepatanLawan > kecepatan {
//				kekuatan -= 6
//				if kekuatan < 3 {
//					kekuatan = 3
//				}
//				terkalahkan++
//			} else {
//				if kekuatan > kecepatan {
//					kekuatan -= 6
//					if kekuatan < 3 {
//						kekuatan = 3
//					}
//					terkalahkan++
//				} else {
//					kecepatan -= 6
//					if kecepatan < 3 {
//						kecepatan = 3
//					}
//					terkalahkan++
//				}
//			}
//		}
//	}
//
//	fmt.Println(terkalahkan, kekuatan, kecepatan)
//}
//
//package main
//
//import "fmt"
//
//func main() {
//	var uangQirat int
//	fmt.Scanln(&uangQirat)
//
//	dinar, dirham, fals, qirat := hitungUang(uangQirat)
//	fmt.Println(dinar, dirham, fals, qirat)
//}
//
//func hitungUang(uangQirat int) (int, int, int, int) {
//	fals := uangQirat / 6
//	sisaQirat := uangQirat % 6
//
//	dirham := fals / 10
//	sisaFals := fals % 10
//
//	dinar := dirham / 10
//	sisaDirham := dirham % 10
//
//	return dinar, sisaDirham, sisaFals, sisaQirat
//}

// package main
//
// import (
//
//	"fmt"
//
// )
//
//	func main() {
//		var tinggi int
//		fmt.Scanln(&tinggi)
//
//		if tinggi%2 == 0 || tinggi <= 0 {
//			fmt.Println("Tinggi harus bilangan asli positif ganjil.")
//			return
//		}
//
//		cetakSegitiga(tinggi)
//	}
//
//	func cetakSegitiga(tinggi int) {
//		for i := 0; i < tinggi; i++ {
//			for j := 0; j <= i; j++ {
//				fmt.Print("*")
//			}
//			fmt.Println()
//		}
//	}
//
// package main
//
// import (
//
//	"fmt"
//	"math"
//
// )
//
//	func main() {
//		// Koordinat posisi pesawat, planet A, dan planet B
//		var xFalcon, yFalcon, zFalcon float64
//		var xA, yA, zA float64
//		var xB, yB, zB float64
//
//		fmt.Scan(&xFalcon, &yFalcon, &zFalcon)
//		fmt.Scan(&xA, &yA, &zA)
//		fmt.Scan(&xB, &yB, &zB)
//
//		jarakA := hitungJarak(xFalcon, yFalcon, zFalcon, xA, yA, zA)
//		jarakB := hitungJarak(xFalcon, yFalcon, zFalcon, xB, yB, zB)
//
//		//fmt.Printf("Jarak ke planet A: %.2f\n", jarakA)
//		//fmt.Printf("Jarak ke planet B: %.2f\n", jarakB)
//
//		if jarakA < jarakB {
//			fmt.Printf("%.2f A", jarakA)
//		} else {
//			fmt.Printf("%.2f B", jarakB)
//		}
//	}
//
//	func hitungJarak(x1, y1, z1, x2, y2, z2 float64) float64 {
//		jarak := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2) + math.Pow(z2-z1, 2))
//		return jarak
//	}
//
//package main
//
//import "fmt"
//
//type State struct {
//	position  int
//	health    int
//	treasures int
//	hasRing   bool
//	alive     bool
//}
//
//func main() {
//	var bilbo, dwarf, orc, smaug int
//	fmt.Scan(&bilbo, &dwarf, &smaug)
//	var route string
//	fmt.Scan(&route)
//
//	state := State{
//		position:  0,
//		health:    bilbo,
//		treasures: 0,
//		hasRing:   false,
//		alive:     true,
//	}
//
//	for _, char := range route {
//		switch char {
//		case '.':
//			if state.alive {
//				state.position++
//			}
//		case 'S':
//			damage := orc
//			if char == 'S' {
//				damage = smaug
//			}
//			state.health -= damage
//			if state.health <= 0 {
//				state.alive = false
//			}
//			if state.alive {
//				state.position++
//			}
//		case 'O':
//			if state.alive {
//				state.health -= dwarf
//				state.position++
//			}
//		case 'H':
//			if state.alive {
//				state.health = bilbo
//				state.treasures++
//				state.position++
//			}
//		case 'A':
//			if state.alive {
//				state.hasRing = true
//				state.position++
//			}
//		}
//	}
//
//	if state.alive {
//		fmt.Println(state.position, state.health, state.treasures, state.hasRing)
//	} else {
//		fmt.Println(state.position, state.health, state.treasures, false)
//	}
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func panjangX(R, T float64) float64 {
//	return R * math.Cos(T*math.Pi/180)
//}
//
//func panjangY(R, T float64) float64 {
//	return R * math.Sin(T*math.Pi/180)
//}
//
//func main() {
//	var R, T float64
//	//fmt.Print("Masukkan panjang garis (R): ")
//	fmt.Scan(&R, &T)
//	//fmt.Print("Masukkan sudut (T) dalam derajat: ")
//	//fmt.Scan(&T)
//
//	panjangX := panjangX(R, T)
//	panjangY := panjangY(R, T)
//	fmt.Printf("%.2f %.2f", panjangX, panjangY)
//
//	//fmt.Printf("Panjang garis dalam sumbu x: %.2f\n", panjangX)
//	//fmt.Printf("Panjang garis dalam sumbu y: %.2f\n", panjangY)
//}

package main

import "fmt"

func reamure(C float64) float64 {
	return 4.0 / 5.0 * C
}

func fahrenheit(C float64) float64 {
	return 9.0/5.0*C + 32
}

func kelvin(C float64) float64 {
	return C + 273
}

func celcius(C float64) float64 {
	return C
}

func main() {
	var startIndex, untilIndex, rangeIndex int

	fmt.Scan(&startIndex, &untilIndex, &rangeIndex)
	//fmt.Println("Celcius \t Reamur \t Fahrenheit \t Kelvin")

	fmt.Printf("%10s %10s %10s %10s\n", "Celcius", "Reamur", "Fahrenheit", "Kelvin")

	for i := startIndex; i <= untilIndex; i += rangeIndex {
		C := float64(i)
		Re := reamure(C)
		Fa := fahrenheit(C)
		Ke := kelvin(C)
		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", C, Re, Fa, Ke)
	}
}
