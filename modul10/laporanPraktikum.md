<h1 align="center">Laporan Praktikum Modul 14 </h1>
<p align="center">Revand Nabil Wicaksana - 109082500002</p>

## Unguided 

### 1. [Soal]

Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya
Hercules sangat suka mengunjungi semua kerabatnya itu.
Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program
rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut
membesar menggunakan algoritma selection sort.
Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat
Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m <1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan
rangkaian bilangan bulat positif, nomor rumah para kerabat.

Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing-
masing daerah.

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul14/output-soal1.png)
Program ini digunakan untuk menampilkan berat terkecil dan berat terbesar dari kelinci dengan menggunakan array untuk menyimpan berat kelincinya

### 2. [Soal]
Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program
ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan
dijual.
Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan
y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya
ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil
yang menyatakan banyaknya ikan yang akan dijual.
Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan
total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang
dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah
sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul14/output-soal2.png)
Program ini digunakan untuk menampilkan total berat ikan yang ada didalam wadah dan menghitung rata rata berat ikan di setiap wadahnya dengan menggunakan array untuk menyimpan berat ikannya.
### 3. [Soal]

Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data
berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data
yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
Buatlah program dengan spesifikasi subprogram sebagai berikut:

type arrBalita [100]float64
func hitungMinMax(arrBerat arrBalita; bMin, bMax *float64) {
/* I.S. Terdefinisi array dinamis arrBerat
Proses: Menghitung berat minimum dan maksimum dalam array
F.S. Menampilkan berat minimum dan maksimum balita */
...
}
function rerata (arrBerat arrBalita) real {
/* menghitung dan mengembalikan rerata berat balita dalam array */
...
}

```go
package main

import "fmt"

const NMAX = 100

type arrBalita [NMAX]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var berat arrBalita
	var n int
	var bMin, bMax, rata float64

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	hitungMinMax(berat, n, &bMin, &bMax)
	rata = rerata(berat, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul14/output-soal3.png)
Program ini digunakan untuk menampilkan berat balita minimum, maksimum dan rerata berat balita pada sebuah posyandu, menggunakan array untuk menyimpan berat balitanya

