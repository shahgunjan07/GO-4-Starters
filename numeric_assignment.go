package main

import (
	"fmt"
)

func main() {
	var x float64
	var y float64

	x, y = 1, 2

	fmt.Printf("\n Value of x=%v , Type of x=%T", x, x)
	fmt.Printf("\n Value of y=%v , Type of y=%T", y, y)

	var mean float64
	mean = (x + y) / 2.0

	fmt.Printf("\n Value of mean=%v , Type of mean=%T", mean, mean)
}
