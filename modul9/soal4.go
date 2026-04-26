package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input string
	*n = 0

	fmt.Print("Text (without space): ")
	fmt.Scan(&input)

	for i := 0; i < len(input); i++ {
		if input[i] == '.' {
			break
		}
		t[*n] = rune(input[i])
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	i := 0
	j := n - 1

	for i < j {
		t[i], t[j] = t[j], t[i]
		i++
		j--
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	fmt.Print("Reverse text: ")
	balikanArray(&tab, m)
	cetakArray(tab, m)
	fmt.Print("Palindrom? ")
	if palindrom(tab, m) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}