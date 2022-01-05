package main

import (
	"fmt"
)


/* An example of pass by value */
func double(val int) {
	fmt.Println("---")
	orig := val
	val *= 2
	fmt.Printf("Double value of %v is %v \n", orig, val)
}

func cleanup(val int) {
	fmt.Println("---")
	fmt.Printf("Resource Cleanup : Resouce id : %v \n", val)
}

func main() {
	/**
	 * defer works in similar manner as "finally" block in JAVA.
	 * You can use defer keyword against methods which are responsible for resource cleanup.
	 * defer method is called at the end, in reverse order.
	 * defer keyword is used for garbage collection purpose.
	 **/
    defer cleanup(1)
	defer cleanup(2)

	a := 2
	double(a)
}