<h1 align="center">Laporan Praktikum Modul 3 </h1>
<p align="center">Revand Nabil Wicaksana - 109082500002</p>

## Unguided 

### 1. [Soal]

Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila
diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)
berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan
koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan
radiusnya.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat
dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik
sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan
bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik
di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

```go
package main

import (
    "fmt"
    "math"
)

type titik struct {
    x, y float64
}

type lingkaran struct {
    pusat  titik
    radius float64
}

func jarak(p, q titik) float64 {
    dx := p.x - q.x
    dy := p.y - q.y
    return math.Sqrt(dx*dx + dy*dy)
}

func didalam(c lingkaran, p titik) bool {
    return jarak(c.pusat, p) <= c.radius
}

func main() {
    var c1, c2 lingkaran
    var p titik

    fmt.Scan(&c1.pusat.x, &c1.pusat.y, &c1.radius)
    fmt.Scan(&c2.pusat.x, &c2.pusat.y, &c2.radius)
    fmt.Scan(&p.x, &p.y)

   	dlm1 := didalam(c1, p)
    dlm2 := didalam(c2, p)

    if   	dlm1 && dlm2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if   	dlm1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if dlm2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul9/output-soal1.png)
Program ini digunakan untuk menentukan sebuah posisi titik terharap 2 lingkaran, program menggunakan rumus jarak untuk menghitung jarak titik ke pusat masing masing lingkaran, lalu program akan memeriksan apakah jarak itu lebih kecil atau sama dengan jari jari lingkaran, jika seperti itu berarti titik berada didalam lingkaran.

### 2. [Soal]
2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program
yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array
memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat
menampilkan beberapa informasi berikut:
a. Menampilkan keseluruhan isi dari array.
b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah
genap).
d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh
dari masukan pengguna.
e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.
Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array
tersebut.
h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi
tersebut.

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul9/output-soal2.png)
Program ini meminta user untuk menginputkan beberapa bilangan bulat yang setelah itu akan disimpan di dalam array, point a menampilkan keseluruhan isi array yang user inputkan, lalu point b menampilkan elemen array dengan index ganjil dengan cara doimodulus jika tidak sama dengan 0 maka index tersebut ganjil, lalu point c menampilkan elemen dengan index genap, jika genap menggunakan modulus jika habis dibagi 2 maka index tersebut genap, lalu d menampilkan elemen array dengan index berkelipatan x dengan x diinputkan oleh user, disini hanya menambahkan inputan berupa bilangan bulat x lalu menggunakan perulangan dan if, lalu e menampilkan untuk menghapus suatu index, disini menambahkan inputan user untuk menginputkakn bilangan bulat tujuannya untuk menghapus sebuah index, lalu f menampilkan rata rata dari elemen arraydengan cara menjumlahkan semua elemen array lalu dibagi dengan jumlah indexnya (dari 1 bukan 0), lalu g menampilkan standar deviasi dengan cara menghitung selisih per elemen yaitu dikurangi dengan rata rata arraynya, kemudian semua selisih itu dijumlahkan lalu hitung varians telebih dahulu dengan cara jumlah selisih tadidibagi dengan index arraynya, baru setelah itu masukan rumus standar deviasi menggunakan math karena akar kuadrat, kemudian yang terakhir h menampilkan frekuensi disini menambahkan inputan untuk mencari frekuensi sebuah elemen di array tersebut.
### 3. [Soal]

Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang
memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang
digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.
Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian
program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan
dalam array adalah nama-nama klub yang menang saja.
Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir
program, tampilkan daftar klub yang memenangkan pertandingan.

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul9/output-soal3.png)
Program ini digunakan untuk menentukan pemenang dari setiap pertandingan yang berjalan dengan menyimpan hasilnya didalam array.

### 4. [Soal]

Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk
membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa
apakah membentuk palindrom.

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input string
	*n = 0

	fmt.Print("Text (without space): ")
	fmt.Scan(&input)

	for i := 0; i < len(input); i++ {
		if input[i] == '.' {
			break
		}
		t[*n] = rune(input[i])
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	i := 0
	j := n - 1

	for i < j {
		t[i], t[j] = t[j], t[i]
		i++
		j--
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	fmt.Print("Reverse text: ")
	balikanArray(&tab, m)
	cetakArray(tab, m)
	fmt.Print("Palindrom? ")
	if palindrom(tab, m) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul9/output-soal4.png)
Program ini digunakan untuk menentukan mereverse suatu huruf yang diinputkan, dan mengecek apakah huruf yang diinputkan itu palindrom atau bukan, dengan menjadikan sebuah pseudocode yang diberikan menjadi kode di vs code, yang awal itu ada const NMAX yang berarti max array hingga 127 inputan saja.
