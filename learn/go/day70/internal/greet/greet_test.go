package greet_test

import (
	"testing"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day70/internal/greet"
)

func TestTitle(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"ada", "Ada"},
		{"  bob  ", "Bob"},
		{"", "Friend"},
	}
	for _, tc := range tests {
		if got := greet.Title(tc.in); got != tc.want {
			t.Fatalf("Title(%q) = %q", tc.in, got)
		}
	}
}

func TestMessage(t *testing.T) {
	if got := greet.Message("go"); got != "Hello, Go!" {
		t.Fatalf("Message = %q", got)
	}
}

func BenchmarkMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		greet.Message("benchmark")
	}
}
