package sliceutil

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3, 4}); got != 10 {
		t.Errorf("Sum = %d, want 10", got)
	}
}

func TestUnique(t *testing.T) {
	got := Unique([]int{1, 2, 2, 3, 1})
	want := []int{1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("Unique len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Unique[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func BenchmarkSum(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(nums)
	}
}

func BenchmarkUnique(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i % 100
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Unique(nums)
	}
}

func ExampleSum() {
	fmt.Println(Sum([]int{1, 2, 3}))
	// Output: 6
}

func ExampleFilterEven() {
	fmt.Println(FilterEven([]int{1, 2, 3, 4, 5}))
	// Output: [2 4]
}
