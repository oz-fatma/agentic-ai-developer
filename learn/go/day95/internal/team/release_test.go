package team

import (
	"strings"
	"testing"
)

func TestReleaseNotes(t *testing.T) {
	notes := ReleaseNotes("cache warmup")
	if !strings.Contains(notes, "v1.0.0") {
		t.Fatal("missing version")
	}
	if !strings.Contains(notes, "cache warmup") {
		t.Fatal("missing feature")
	}
	if !strings.Contains(notes, "Review checklist:") {
		t.Fatal("missing checklist header")
	}
}
