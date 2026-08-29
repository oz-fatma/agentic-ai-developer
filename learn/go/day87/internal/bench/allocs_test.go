package bench

import "testing"

func BenchmarkConcat(b *testing.B) {
	for b.Loop() {
		_ = Concat(128)
	}
}

func BenchmarkBuilder(b *testing.B) {
	for b.Loop() {
		_ = Builder(128)
	}
}

func BenchmarkAllocSlice(b *testing.B) {
	for b.Loop() {
		_ = AllocSlice(256)
	}
}
