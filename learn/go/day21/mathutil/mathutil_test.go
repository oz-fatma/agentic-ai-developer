package mathutil

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	if got != 5 {
		t.Errorf("Add(2, 3) = %d, want 5", got)
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(4, 5)
	if got != 20 {
		t.Fatalf("Multiply(4, 5) = %d, want 20", got)
	}
}

func TestIsEven(t *testing.T) {
	if !IsEven(4) {
		t.Error("IsEven(4) = false, want true")
	}
	if IsEven(3) {
		t.Error("IsEven(3) = true, want false")
	}
}

func TestAbs(t *testing.T) {
	if Abs(-7) != 7 {
		t.Errorf("Abs(-7) = %d, want 7", Abs(-7))
	}
	if Abs(5) != 5 {
		t.Errorf("Abs(5) = %d, want 5", Abs(5))
	}
}
