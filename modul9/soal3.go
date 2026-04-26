package main

import "fmt"

func main() {
	var kA, kB string
	var s1, s2 int
	var hasil [1000]string
	pertandingan := 0

	fmt.Print("Klub A : ")
	fmt.Scan(&kA)
	fmt.Print("Klub B : ")
	fmt.Scan(&kB)

	for {
		fmt.Printf("Pertandingan %d : ", pertandingan+1)
		fmt.Scan(&s1, &s2)

		if s1 < 0 || s2 < 0 {
			break
		}
		if s1 > s2 {
			hasil[pertandingan] = kA
		} else if s2 > s1 {
			hasil[pertandingan] = kB
		} else {
			hasil[pertandingan] = "Draw"
		}
		pertandingan++
		}
	for i := 0 ; i < pertandingan ; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil[i])
	}
	fmt.Print("Pertandingan selesai")
}