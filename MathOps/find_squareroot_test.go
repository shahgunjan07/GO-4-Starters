package MathOps

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"testing"
)

func almostEqual(value1 float64, value2 float64) bool {
	return math.Abs(value1-value2) <= 0.001
}

func TestMany(t *testing.T) {
	file, err := os.Open("sqrt_input.csv")

	if err != nil {
		t.Fatalf("Can not open test data file %s", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatalf("Error generated while reading test data file : %s", err)
		}

		valueAsInput, err := strconv.ParseFloat(record[0], 64)

		if err != nil {
			t.Fatalf("Bad Input : %s", err)
		}

		expectedResultAsOutput, err := strconv.ParseFloat(record[1], 64)

		if err != nil {
			t.Fatalf("Bad Input : %s", err)
		}

		t.Run(fmt.Sprintf("%f", valueAsInput), func(t *testing.T) {

			result, err := findSquareRoot(valueAsInput)

			if err != nil {
				t.Fatalf("Error occured while trying to calculate square root : %s", err)
			}

			if !almostEqual(result, expectedResultAsOutput) {
				t.Fatalf("Test execution failed ! Expected Value : %f != Actual Ouput : %f", expectedResultAsOutput, result)
			}
		})

	}
}
