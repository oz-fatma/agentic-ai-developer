package service

import "testing"

func TestSimpleGreeter(t *testing.T) {
	g := NewSimpleGreeter("Hello")
	got := g.Greet("Go")
	want := "Hello, Go!"
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

type mockGreeter struct {
	called bool
}

func (m *mockGreeter) Greet(name string) string {
	m.called = true
	return "mock:" + name
}

func TestGreeterInterface(t *testing.T) {
	var g Greeter = &mockGreeter{}
	if g.Greet("x") != "mock:x" {
		t.Fatal("mock greeter failed")
	}
}
