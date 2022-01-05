package main

import (
	"fmt"
)

func main() {
	x := 1.0
	y := 2.0

	fmt.Printf("\n Value of x=%v , Type of x=%T", x, x)
	fmt.Printf("\n Value of y=%v , Type of y=%T", y, y)

	mean := (x + y) / 2.0

	fmt.Printf("\n Value of mean=%v , Type of mean=%T", mean, mean)
}
