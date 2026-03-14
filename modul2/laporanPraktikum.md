# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Revand Nabil Wicaksana] - [109082500002]</p>

## Unguided 

### 1. [Soal]
#### soal1.go
Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan
masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang
diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?

```go
package main

import "fmt"

func main() {
var (
satu, dua, tiga string
temp string
)
fmt.Print("Masukan input string: ")
fmt.Scanln(&satu)
fmt.Print("Masukan input string: ")
fmt.Scanln(&dua)
fmt.Print("Masukan input string: ")
fmt.Scanln(&tiga)
fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
temp = satu
satu = dua
dua = tiga
tiga = temp
fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul2/output-soal1.png)
program diatas bertujuan untuk mengubah urutan bilangan yang diinputkan menggunakan variabel temp, kegunaan variabel temp ini yaitu untuk digunakan menyimpan inputan pertama, nanti urutannya akan bergeser ke urutan kedua pindah ke urutan pertama, urutan ketiga pindah ke urutan kedua, urutan pertama yang disimpan didalam temp pindah ke urutan ketiga


### 2. [Soal]
Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan
praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana
susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta
untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan
warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’,
‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.
Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi
sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan
warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false
untuk urutan warna lainnya.

```go
package main

import "fmt"

func main() {
	var ur1, ur2, ur3, ur4 string
	berhasil := true

	for a := 1; a <= 5; a++ {
		fmt.Print("Percobaan ", a, ": ")
		fmt.Scan(&ur1, &ur2, &ur3, &ur4)

		if !(ur1 == "merah" && ur2 == "kuning" && ur3 == "hijau" && ur4 == "ungu") {
			berhasil = false
		}
	}
	fmt.Println("BERHASIL:", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul2/output-soal2.png)
program diatas bertujuan untuk menilai bahwa hasil dari zat cair itu konsisten dari percobaan pertama hingga terakhir, percobaan ke 1-5 menggunakan perulangan, lalu untuk menentukan true or falsenya menggunakan  if seperti dicodingan saya.


### 3. [Soal]
#### soal3.go
PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka,
buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan
sebagai berikut!
Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam
gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500
gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500
gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang
kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.

```go
package main

import "fmt"

func main() {

	var berat, biayaKg, biayaSisa, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	sisa := berat % 1000

	biayaKg = kg * 10000

	if kg > 10 {
		biayaSisa = 0
	} else {
		if sisa >= 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}
	}

	total = biayaKg + biayaSisa

	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya: Rp.", total)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/revandnabil/109082500002-Revand-Nabil-Wicaksana/modul2/output-soal3.png)
program diatas bertujuan untuk menghitung biaya pengiriman PT POS, karena inputannya berupa gram maka kita harus konverensikan menjadi kilogram terlebih dahulu dengan menggunakan rumus berat dibagi 1000, lalu menghitung sisa gramnya menggunakan modulus 1000, lalu kemudian ada rumus untuk menghitung biaya pengiriman berdasarkan kilogram. Jika berat lebih dari 10 kg maka sisa gram tidak dikenakan biaya atau Rp. 0, lalu jika berat tidak lebih dari 10kg maka (jika sisa gramnya diatas 499 maka ada biaya tambahan sebesar Rp. 5 / gramnya, jika sisa gramnya dibawah 500 maka biaya tambahannya Rp. 15 / gramnya), lalu menghitung total biayanya dengan cara menambahkan biaya kilogram dengan biaya sisanya.
