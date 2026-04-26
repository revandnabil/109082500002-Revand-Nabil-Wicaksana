package main

import (
	"fmt"
	"math"
)

func main() {
	var n [5]int
// a.
	fmt.Println("a.Menampilkan keseluruhan isi dari array: ")
	for i := 0 ; i < len(n) ; i++ {
		fmt.Scan(&n[i])
	}
// b.
	fmt.Println()
	fmt.Println("b.Menampilkan elemen-elemen array dengan indeks ganjil saja: ")
	for i := 0 ; i < len(n) ; i++ {
		if i%2 != 0 {
		fmt.Print(n[i], " ")
	}
}
// c.
	fmt.Println()
	fmt.Println("c.Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap): ")
	for i := 0 ; i < len(n) ; i++ {
		if i%2 == 0 {
		fmt.Print(n[i], " ")
		}
}
// d.
	var x int 
	fmt.Println()
	fmt.Println("d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x.x bisa diperoleh dari masukan pengguna.")
	fmt.Print("Masukan x: ")
	fmt.Scan(&x)
	for i := 0 ; i < len(n) ; i++ {
		if i%x == 0 {
		fmt.Print(n[i], " ")
		}
}
// e.
	var hapus int
	fmt.Println()
	fmt.Println("e.Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil.")
	fmt.Print("Masukan index ke berapa yang ingin dihapus: ")
	fmt.Scan(&hapus)
	if hapus >= 0 && hapus < len(n) {
		for i := 0; i < len(n); i++ {
			if i != hapus {
				fmt.Print(n[i], " ")
			}
		}
	} else {
		fmt.Println("-")
	}
// f.
	var ratarata float64
	var jumlah int
	fmt.Println()
	fmt.Println("f. Menampilkan rata-rata dari bilangan yang ada di dalam array: ")
	for i := 0 ; i < len(n) ; i++ {
		jumlah = jumlah + n[i] 
	}
	ratarata = float64(jumlah) / float64(len(n))
	fmt.Print(ratarata)
// g.
	var totalSelisih float64
	fmt.Println()
	fmt.Println("g.Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.")
	for i := 0 ; i < len(n) ; i++ {
		selisih := float64(n[i]) - ratarata
		totalSelisih += selisih * selisih
}
	varians := totalSelisih / float64(len(n))
	standarDeviasi := math.Sqrt(varians)
	fmt.Printf("%.2f\n", standarDeviasi)
// h.
	var a int 
	frekuensi := 0
	fmt.Println()
	fmt.Println("h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.")
	fmt.Print("Input bilangan: ")
	fmt.Scan(&a)
	for i := 0 ; i < len(n) ; i++ {
		if a == n[i] {
			frekuensi++
		}
}
	fmt.Print(frekuensi)
}