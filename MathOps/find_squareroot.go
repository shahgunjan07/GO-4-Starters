package MathOps

import (
	"fmt"
	"math"
)

/**
 * This is an example of function which returns two values.
 * value1 : result of square root
 * value2 : error
 */
func findSquareRoot(val float64) (float64, error) {
	if val < 0 {
		return 0.0, fmt.Errorf("Square root of negative number is not possible")
	} else {
		return math.Sqrt(val), nil
	}
}
