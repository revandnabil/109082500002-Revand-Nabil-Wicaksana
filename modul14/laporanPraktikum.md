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
Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan
rangkaian bilangan bulat positif, nomor rumah para kerabat.
Keluaran terdiri dari n baris, yaitu

```go
package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main() {
	var n, m int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSort(rumah)

		for j := 0; j < m; j++ {
			fmt.Print(rumah[j])
			if j < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul14/output-soal1.png)
Program ini digunakan untuk mengurutkan nomer rumah dari inputan terkecil hingga terbesar dari rumah Hercules.

### 2. [Soal]
Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu
diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena
nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah
program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil
lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor
genap terurut mengecil.
Format Masukan masih persis sama seperti sebelumnya.
Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk
nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.
Keterangan: Terdapat 3 daerah dalam contoh masukan. Baris kedua berisi campuran
bilangan ganjil dan genap. Baris berikutnya hanya berisi bilangan ganjil, dan baris terakhir
hanya berisi bilangan genap.
Petunjuk:
• Waktu pembacaan data, bilangan ganjil dan genap dipisahkan ke dalam dua array
yang berbeda, untuk kemudian masing-masing diurutkan tersendiri.
• Atau, tetap disimpan dalam satu array, diurutkan secara keseluruhan. Tetapi pada
waktu pencetakan, mulai dengan mencetak semua nilai ganjil lebih dulu, kemudian
setelah selesai cetaklah semua nilai genapnya.

```go
package main

import "fmt"

func sortAscending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func sortDescending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[idxMax] {
				idxMax = j
			}
		}
		arr[i], arr[idxMax] = arr[idxMax], arr[i]
	}
}

func main() {
	var n, m int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&m); err != nil {
			return
		}

		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {
			var val int
			fmt.Scan(&val)
			if val%2 != 0 {
				ganjil = append(ganjil, val)
			} else {
				genap = append(genap, val)
			}
		}

		sortAscending(ganjil)
		sortDescending(genap)
		hasil := append(ganjil, genap...)

		for j := 0; j < len(hasil); j++ {
			fmt.Print(hasil[j])
			if j < len(hasil)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul14/output-soal2.png)
Program ini digunakan untuk mengurutkan nomer rumah dari ganjil terkecil lalu genap dari yang terbesar, karena rumah kerabat Hecules di sebelah kiri jalan nomernya ganjil semua, sementara di seberang kanan jalaln nomernya genap semua.

### 3. [Soal]

Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan
tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan
sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu
problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba
untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?
"Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data
genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua
data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke
bawah."
Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah
terbaca, jika data yang dibaca saat itu adalah 0.
Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000
data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak
termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313.
Keluaran adalah median yang diminta, satu data per baris.
Keterangan:
Sampai bilangan 0 yang pertama, data terbaca adalah 7 23 11, setelah tersusun: 7 11 23,
maka median saat itu adalah 11.

Halaman 100 | M o d u l P r a k t i k u m A l g o r i t m a d a n P e m r o g r a m a n 2
Sampai bilangan 0 yang kedua, data adalah 7 23 11 5 19 2 29 3 13 17, setelah tersusun
diperoleh: 2 3 5 7 11 13 17 19 23 29. Karena ada 10 data, genap, maka median adalah
(11+13)/2=12.
Petunjuk:
Untuk setiap data bukan 0 (dan bukan marker -5313541) simpan ke dalam array, Dan setiap
kali menemukan bilangan 0, urutkanlah data yang sudah tersimpan dengan menggunakan
metode insertion sort dan ambil mediannya.
...


```go
package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func main() {
	var val int
	var data []int

	for {
		fmt.Scan(&val)
		if val == -5313541 {
			break
		}

		if val == 0 {
			insertionSort(data)

			n := len(data)
			if n > 0 {
				if n%2 != 0 {
					median := data[n/2]
					fmt.Println(median)
				} else {
					median := (data[(n/2)-1] + data[n/2]) / 2
					fmt.Println(median)
				}
			}
		} else {
			data = append(data, val)
		}
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul14/output-soal3.png)
Program ini digunakan untuk menghitung nilai median dari sekumpulan angka yang diinput secara bertahap.

