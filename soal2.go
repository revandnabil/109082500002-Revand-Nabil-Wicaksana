package main

import "fmt"

func main() {
	var inputan int
	var n int

	fmt.Scan(&inputan)

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