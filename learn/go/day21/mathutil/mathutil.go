package mathutil

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func IsEven(n int) bool {
	return n%2 == 0
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
