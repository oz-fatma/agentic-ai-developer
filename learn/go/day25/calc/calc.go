package calc

import "fmt"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func Pow(base, exp int) int {
	if exp < 0 {
		return 0
	}
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
