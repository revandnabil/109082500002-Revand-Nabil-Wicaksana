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
