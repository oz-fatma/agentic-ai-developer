package mathutil

import "errors"

var ErrDivideByZero = errors.New("divide by zero")

func Add(a, b int) int { return a + b }

func Subtract(a, b int) int { return a - b }

func Multiply(a, b int) int { return a * b }

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Max(values ...int) int {
	if len(values) == 0 {
		return 0
	}
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}
	return m
}
