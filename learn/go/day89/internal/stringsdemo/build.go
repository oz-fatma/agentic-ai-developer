package stringsdemo

import (
	"strings"
	"time"
)

func Concat(parts []string) (string, time.Duration) {
	start := time.Now()
	s := ""
	for _, p := range parts {
		s += p
	}
	return s, time.Since(start)
}

func WithBuilder(parts []string) (string, time.Duration) {
	start := time.Now()
	var b strings.Builder
	for _, p := range parts {
		b.WriteString(p)
	}
	return b.String(), time.Since(start)
}
