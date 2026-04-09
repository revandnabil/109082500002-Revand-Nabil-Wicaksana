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