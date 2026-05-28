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