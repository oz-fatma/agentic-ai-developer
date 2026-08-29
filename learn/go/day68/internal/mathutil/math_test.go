package mathutil_test

import (
	"errors"
	"testing"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day68/internal/mathutil"
)

func TestAdd(t *testing.T) {
	if got := mathutil.Add(2, 3); got != 5 {
		t.Fatalf("Add = %d", got)
	}
}

func TestSubtract(t *testing.T) {
	if got := mathutil.Subtract(10, 4); got != 6 {
		t.Fatalf("Subtract = %d", got)
	}
}

func TestMultiply(t *testing.T) {
	if got := mathutil.Multiply(3, 7); got != 21 {
		t.Fatalf("Multiply = %d", got)
	}
}

func TestDivide(t *testing.T) {
	got, err := mathutil.Divide(10, 2)
	if err != nil || got != 5 {
		t.Fatalf("Divide = %d err = %v", got, err)
	}
	_, err = mathutil.Divide(1, 0)
	if !errors.Is(err, mathutil.ErrDivideByZero) {
		t.Fatalf("Divide by zero err = %v", err)
	}
}

func TestAbs(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{0, 0},
		{5, 5},
		{-5, 5},
	}
	for _, tc := range cases {
		if got := mathutil.Abs(tc.in); got != tc.want {
			t.Fatalf("Abs(%d) = %d", tc.in, got)
		}
	}
}

func TestMax(t *testing.T) {
	if got := mathutil.Max(1, 9, 3); got != 9 {
		t.Fatalf("Max = %d", got)
	}
}
