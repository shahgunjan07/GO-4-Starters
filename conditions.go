package main

import (
	"fmt"
)

func main() {

	x := 2

	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknown Number")
	}

	switch {
	case x > 100:
		fmt.Println("x is very big")
	case x > 2:
		fmt.Println("x is big")
	default:
		fmt.Println("x is small")
	}
}
