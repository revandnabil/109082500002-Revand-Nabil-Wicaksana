package main

import (
    "fmt"
    "math"
)

type titik struct {
    x, y float64
}

type lingkaran struct {
    pusat  titik
    radius float64
}

func jarak(p, q titik) float64 {
    dx := p.x - q.x
    dy := p.y - q.y
    return math.Sqrt(dx*dx + dy*dy)
}

func didalam(c lingkaran, p titik) bool {
    return jarak(c.pusat, p) <= c.radius
}

func main() {
    var c1, c2 lingkaran
    var p titik

    fmt.Scan(&c1.pusat.x, &c1.pusat.y, &c1.radius)
    fmt.Scan(&c2.pusat.x, &c2.pusat.y, &c2.radius)
    fmt.Scan(&p.x, &p.y)

   	dlm1 := didalam(c1, p)
    dlm2 := didalam(c2, p)

    if   	dlm1 && dlm2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if   	dlm1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if dlm2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}