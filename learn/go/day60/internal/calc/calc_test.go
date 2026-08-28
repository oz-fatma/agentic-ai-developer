package calc

import "testing"

func TestSumRange(t *testing.T) {
	if got := SumRange(1, 5); got != 15 {
		t.Fatalf("SumRange(1,5)=%d want 15", got)
	}
	if got := SumRange(5, 1); got != 15 {
		t.Fatalf("reversed range=%d want 15", got)
	}
}

func TestAverage(t *testing.T) {
	if got := Average(10, 20, 30); got != 20 {
		t.Fatalf("Average=%v want 20", got)
	}
	if got := Average(); got != 0 {
		t.Fatalf("empty average=%v want 0", got)
	}
}

func TestIsPrime(t *testing.T) {
	cases := map[int]bool{
		0: false, 1: false, 2: true, 3: true, 4: false, 17: true, 18: false,
	}
	for n, want := range cases {
		if got := IsPrime(n); got != want {
			t.Fatalf("IsPrime(%d)=%v want %v", n, got, want)
		}
	}
}
