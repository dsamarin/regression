package main

import (
	"fmt"
)

type Dataset struct {
	n int
	avgXY float64
	avgX float64
	avgY float64
	avgX2 float64
}

func (d *Dataset) Add(x float64, y float64) {
	fmt.Print("Add: {", x, ", ", y, "}\n")

	if d.n == 0 {
		d.avgXY = x * y
		d.avgX = x
		d.avgY = y
		d.avgX2 = x * x
	} else {
		nn := float64(d.n + 1)
		fac := float64(d.n) / nn
		d.avgXY = fac * d.avgXY + (x*y)/nn
		d.avgX  = fac * d.avgX  + (x) / nn
		d.avgY  = fac * d.avgY  + (y) / nn
		d.avgX2 = fac * d.avgX2 + (x*x)/nn
	}

	d.n++
}

func (d *Dataset) SlopeIntercept() (float64, float64) {
	m := (d.avgXY - (d.avgX * d.avgY)) /
		(d.avgX2 - d.avgX * d.avgX)
	b := d.avgY - m * d.avgX
	return m, b
}

func main() {

	// y = mx + b
	secretM := 2.45
	secretB := 19.234

	dataset := Dataset{}
	for i := 1.0; i <= 10.0; i += 1.0 {
		dataset.Add(i, secretM * i + secretB)
	}

	fmt.Println(dataset.SlopeIntercept())

	fmt.Println("Hello, world")
}
