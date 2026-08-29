package perf

import "testing"

func TestSummarize(t *testing.T) {
	got := Summarize([]string{"alpha", "beta", "gamma"}, 2)
	want := "ALPHA\nBETA\nGAMMA\n"
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestSummarizeSingleWorker(t *testing.T) {
	got := Summarize([]string{"one"}, 0)
	if got != "ONE\n" {
		t.Fatalf("got %q", got)
	}
}
