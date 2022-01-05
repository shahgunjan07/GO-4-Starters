package main

import (
	"fmt"
	"math"
)

/**
 * This is an example of function which returns two values.
 * value1 : result of square root
 * value2 : error
*/
func sqrt(val float64) (float64, error) {
	if val < 0 {
		return 0.0, fmt.Errorf("Square root of negative number is not possible")
	} else {
		return math.Sqrt(val), nil
	}
}

/* An example of pass by value */
func double(val int) {
	fmt.Println("---")
	orig := val
	val *= 2
	fmt.Printf("Double value of %v is %v \n", orig, val)
}

/* An example of pass by reference */
func doubleRef(val *int) {
	fmt.Println("---")
	orig := *val
	*val *= 2
	fmt.Printf("Double value of %v is %v \n", orig, *val)
}

func main() {
	// Check sqaure root of non-negative number
	res1, err := sqrt(2)

	if err != nil {
		fmt.Printf("Error : %s \n", err)
	} else {
		fmt.Println(res1)
	}

	// Check sqaure root of negative number
	res2, err := sqrt(-1)

	if err != nil {
		fmt.Printf("Error : %s \n", err)
	} else {
		fmt.Println(res2)
	}

	// Test pass by value
	a := 2
	double(a)
	fmt.Printf("Original value of a after calling double : %v \n", a)

	// Test pass by reference
	doubleRef(&a)
	fmt.Printf("Original value of a after calling double : %v \n", a)
}