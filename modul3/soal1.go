package main

import "fmt"

func factorial(in int, hasil *int) {
	*hasil = 1
	for i := 2; i <= in; i++ {
		*hasil *= i
	}
}
func permutation(n int, r int, hasil *int) {
	var nilaiN, nilaiSisa int
	factorial(n, &nilaiN)
	factorial(n-r, &nilaiSisa)
	*hasil = nilaiN / nilaiSisa
}
func combination(n int, r int, hasil *int) {
	var nilaiN, nilaiR, nilaiSisa int
	factorial(n, &nilaiN)
	factorial(r, &nilaiR)
	factorial(n-r, &nilaiSisa)
	*hasil = nilaiN / (nilaiR * nilaiSisa)
}
func main() {
	var x, y, m, n int
	var perm1, komb1, perm2, komb2 int

	fmt.Scan(&x, &y, &m, &n)

	permutation(x, m, &perm1)
	combination(x, m, &komb1)
	permutation(y, n, &perm2)
	combination(y, n, &komb2)
	fmt.Println(perm1, komb1)
	fmt.Println(perm2, komb2)
}