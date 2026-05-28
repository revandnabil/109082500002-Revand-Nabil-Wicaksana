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