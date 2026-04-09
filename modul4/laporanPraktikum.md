<h1 align="center">Laporan Praktikum Modul 4 </h1>
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
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul4/output-soal1.png)
Program ini digunakan untuk menghitung nilai permutasi dan kombinasi dari input bilangan yang diberikan. Program menerima empat angka yaitu a, b, c, dan d, lalu menghitung permutasi dan kombinasi a terhadap c serta b terhadap d. Perhitungan dilakukan menggunakan fungsi faktorial sebagai dasar rumusnya. Hasilnya ditampilkan dalam dua baris sesuai pasangan perhitungannya.

### 2. [Soal]
Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal
yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan
soal paling banyak dalam waktu paling singkat adalah pemenangnya.
Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program
harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total
soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal.
Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca
di dalam prosedur.
prosedure hitungSkor(in/out soal, skor : integer)
Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah
8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal.
Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan
dalam waktu 5 jam 1 menit (301 menit).
Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang
diperoleh. Nilai adalah total waktu

```go
package main

import "fmt"

func prosesNilai(data [8]int, jml *int, totalWaktu *int) {
	*jml = 0
	*totalWaktu = 0

	for i := 0; i < 8; i++ {
		if data[i] <= 300 {
			*jml++
			*totalWaktu += data[i]
		}
	}
}

func main() {
	var peserta string

	var namaJuara string
	var soalTerbanyak int = -1
	var waktuTerkecil int = 999999
	var jumlahSoal, total int
	var nilai [8]int
	fmt.Scan(&peserta)

	for peserta != "Selesai" {

		for i := 0; i < 8; i++ {
			fmt.Scan(&nilai[i])
		}
		prosesNilai(nilai, &jumlahSoal, &total)

		if jumlahSoal > soalTerbanyak || (jumlahSoal == soalTerbanyak && total < waktuTerkecil) {
			namaJuara = peserta
			soalTerbanyak = jumlahSoal
			waktuTerkecil = total
		}

		fmt.Scan(&peserta)
	}

	fmt.Println(namaJuara, soalTerbanyak, waktuTerkecil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul4/output-soal2.png)
Program ini digunakan untuk menentukan pemenang dari beberapa peserta sesuai dengan jumlah soal yang berhasil diselesaikan dan total waktu. Setiap peserta memasukkan nama dan 8 nilai waktu, lalu program membaca data tersebut sampai input "Selesai".
Fungsi prosesNilai menghitung jumlah soal yang selesai (≤300 menit) dan menjumlahkan total waktunya.
Pemenang dipilih dari soal terbanyak, dan jika sama maka yang waktunya paling kecil akan ditampilkan sebagai hasil akhir.

