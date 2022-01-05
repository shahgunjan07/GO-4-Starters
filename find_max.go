package main


import (
	"fmt"
)

func main() {
	nums := []int{16,8,42,4,23,15}
	max := nums[0]
	resultPosition := 0

	for index, value := range nums[0:] {
		if(value > max) {
			max = value
			resultPosition = index
		}
	}

	fmt.Printf("Max Value is : %v at Index : %v",max, resultPosition)
}
