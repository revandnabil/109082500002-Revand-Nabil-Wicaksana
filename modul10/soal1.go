package main

import "fmt"

func main() {
    const NMAX = 1000
	var berat [NMAX]float64
    var N int

    fmt.Print("Masukan banyak kelinci: ")
    fmt.Scan(&N)

	if N > 0 && N < NMAX {
        for i := 0; i < N; i++ {
            fmt.Printf("Berat kelinci ke-%d: ", i+1)
            fmt.Scan(&berat[i])
        }
    } else {
        fmt.Println("Hitung ayam minimal 1 atau maximal 1000")
    }

    min := berat[0]
    max := berat[0]

    for i := 1; i < N; i++ {
        if berat[i] < min {
            min = berat[i]
        } else if berat[i] > max {
            max = berat[i]
        }
    }
    fmt.Println("Berat kelinci terkecil: ", min)
    fmt.Println("Berat kelinci terbesar: ", max)
}