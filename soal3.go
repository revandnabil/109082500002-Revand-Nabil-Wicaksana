package main

import "fmt"

func hitung(inputan int) {
	var n int
	if inputan%2 != 0 {
	for i := 0; i <= inputan ; i++ {
		n = inputan - i
		fmt.Println("-", n, n)
	}
} else {
	for i := 0; i <= inputan ; i++ {
		n = inputan - i
		fmt.Print("-", n)
}
}
}
func main() {
	var inputan int

	hitung(inputan)
}