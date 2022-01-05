package main

import (
	"fmt"
)

func main() {

	for index := 1; index <= 15; index++ {

		if index%3 == 0 && index%5 == 0 {
			fmt.Println("Fizz Buzz")
			continue
		} else if index%3 == 0 {
			fmt.Println("Fizz")
			continue
		} else if index%5 == 0 {
			fmt.Println("Buzz")
			continue
		} else {
			fmt.Println(index)
		}

	}
}
