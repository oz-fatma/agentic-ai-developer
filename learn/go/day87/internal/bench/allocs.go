package bench

import "strings"

func Concat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "x"
	}
	return s
}

func Builder(n int) string {
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < n; i++ {
		b.WriteByte('x')
	}
	return b.String()
}

func AllocSlice(n int) []byte {
	out := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, byte(i%256))
	}
	return out
}
