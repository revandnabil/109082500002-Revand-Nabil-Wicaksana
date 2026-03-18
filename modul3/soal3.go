package main

import "fmt"

func dalamLingkaran(x, y, cx, cy, r int) bool {
	return (x-cx)*(x-cx)+(y-cy)*(y-cy) <= r*r
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	d1 := dalamLingkaran(x, y, cx1, cy1, r1)
	d2 := dalamLingkaran(x, y, cx2, cy2, r2)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}