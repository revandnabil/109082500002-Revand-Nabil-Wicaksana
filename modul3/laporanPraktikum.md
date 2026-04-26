<h1 align="center">Laporan Praktikum Modul 3 </h1>
<p align="center">Revand Nabil Wicaksana - 109082500002</p>

## Unguided 

### 1. [Soal]

Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika
diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng
untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian
membantu Jonas? (tidak tentunya ya :p)
Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi,
dengan syarat a ≥ c dan b ≥ d.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a
terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan
menggunakan persamaan berikut!
P(n, r) =
n!
(n−r)!
, sedangkan C(n, r) =
n!
r!(n−r)!

```go
package main

import "fmt"

func faktorial(n int) int {
    hasil := 1
    for i := 1; i <= n; i++ {
        hasil *= i
    }
    return hasil
}

func permutasi(n, r int) int {
    return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
    return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)

    fmt.Println(permutasi(a, c), kombinasi(a, c))
    
    fmt.Println(kombinasi(b, d), kombinasi(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul3/output-soal1.png)
Program ini digunakan untuk menghitung nilai permutasi dan kombinasi dari input bilangan yang diberikan. Program menerima empat angka yaitu a, b, c, dan d, lalu menghitung permutasi dan kombinasi a terhadap c serta b terhadap d. Perhitungan dilakukan menggunakan fungsi faktorial sebagai dasar rumusnya. Hasilnya ditampilkan dalam dua baris sesuai pasangan perhitungannya

### 2. [Soal]
2. Diberikan tiga buah fungsi matematika yaitu f (x) = x'2, g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x)
dalam bentuk function.
Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi.
Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b),
dan baris ketiga adalah (hofog)(c)!

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}
func g(x int) int {
	return x - 2
}
func h(x int) int {
	return x + 1
}
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul3/output-soal2.png)
Program diatas digunakan untuk menhitung hasil komposisi. Perhitungan dilakukan dengan cara memasukkan nilai ke fungsi paling dalam, lalu ke fungsi berikutnya hingga selesai. Lalu hasil akhirnya ditampilkan dalam tiga baris sesuai urutan komposisinya.

### 3. [Soal]

Lingkaran Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius
r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)
berdasarkan dua lingkaran tersebut.
Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat
dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik
sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan
bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik
di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

```go
package main

import "fmt"

func dalamLingkaran(x, y, cx, cy, r int) bool {
	return (x-cx)*(x-cx)+(y-cy)*(y-cy) <= r*r
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	d1 := dalamLingkaran(x, y, cx1, cy1, r1)
	d2 := dalamLingkaran(x, y, cx2, cy2, r2)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul3/output-soal3.png)
Program ini digunakan untuk menentukan posisi sebuah titik terhadap dua lingkaran. Setiap lingkaran memiliki titik pusat dan jari-jari, lalu dicek apakah titik berada di dalam atau di luar lingkaran menggunakan rumus jarak. Outputnya berupa keterangan apakah titik berada di dalam lingkaran 1, lingkaran 2, keduanya, atau di luar keduanya.
