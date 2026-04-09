<h1 align="center">Laporan Praktikum Modul 5 </h1>
<p align="center">Revand Nabil Wicaksana - 109082500002</p>

## Unguided 

### 1. [Soal]

Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai
suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat
diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku
ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci
tersebut.

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul5/output-soal1.png)
Program ini pakai rekursif untuk menghitung deret Fibonacci, dimana setiap angka adalah hasil penjumlahan dua angka sebelumnya.

### 2. [Soal]
Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan
menggunakan fungsi rekursif. N adalah masukan dari user.

```go
package main

import "fmt"

func bintang(n int, i int) {
	if i > n {
		return
	}

	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	bintang(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	bintang(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul5/output-soal2.png)
Program diatas untuk mencetak pola bintang menggunakan rekursif. Fungsi akan mencetak bintang sebanyak nilai saat ini, lalu memanggil dirinya lagi untuk baris berikutnya. Proses ini terus berjalan sampai jumlah baris mencapai n bilangan yang diinputkan

### 3. [Soal]
Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari
suatu N, atau bilangan yang apa saja yang habis membagi N.
Masukan terdiri dari sebuah bilangan bulat positif N.
Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).

```go
package main

import "fmt"

func faktor(n int, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Print(i, " ")
	}

	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	faktor(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul5/output-soal3.png)
Program ini mencari semua angka yang bisa membagi habis suatu bilangan n yang diinputkan. Dicek satu-satu dari 1 sampai n, kalau tidak ada sisa berarti itu faktor dan langsung ditampilkan di outputnya. 

