// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {

	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad input: %f, %f (division by zero is undefined)", a, b)
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input: %f (square root of negative number is undefined)", a)
	}
	return math.Sqrt(a), nil
}

func AddMany(inputs ...float64) float64 {
	total := 0.0
	for _, input := range inputs {
		total = Add(total, input)
	}
	return total
}

func DivideMany(inputs ...float64) (float64, error) {
	result := 0.0
	var err error
	for _, input := range inputs {
		if result == 0.0 {
			result = input
		}
		result, err = Divide(result, input)
		if err != nil {
			return 0, fmt.Errorf("bad input: %f can't not be 0", input)
		}
	}
	return result, nil
}
