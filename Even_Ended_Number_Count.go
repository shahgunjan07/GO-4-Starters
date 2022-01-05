package main

import (
	"fmt"
)

func main() {

	count := 0

	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			temp := fmt.Sprintf("%d", i*j)

			if temp[0] == temp[len(temp)-1] {
				count++
				//fmt.Println("Even Ended Number : ", temp)
			}
		}
	}

	fmt.Println("Total Count of Even Ended Number is :", count)
}
