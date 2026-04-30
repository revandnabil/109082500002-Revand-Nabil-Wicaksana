package main

import "fmt"

func main() {
	const NMAX = 1000
	var ikan [NMAX]float64
	var x, y int

	fmt.Print("Banyak ikan yang akan dijual: ")
	fmt.Scan(&x)
	fmt.Print("Kapasitas per wadah: ")
	fmt.Scan(&y)

	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	var totalWadah [NMAX]float64
	var rataWadah [NMAX]float64
	var jumlahWadah int

	for i := 0; i < x; i += y {
		var total float64
		var jumlahIkan int

		for j := i; j < i+y && j < x; j++ {
			total += ikan[j]
			jumlahIkan++
		}

		totalWadah[jumlahWadah] = total
		rataWadah[jumlahWadah] = total / float64(jumlahIkan)
		jumlahWadah++
	}

	fmt.Print("Total berat tiap wadah: ")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("%.2f ", totalWadah[i])
	}

	fmt.Print("\nRata-rata berat tiap wadah: ")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("%.2f ", rataWadah[i])
	}

	fmt.Println()
}