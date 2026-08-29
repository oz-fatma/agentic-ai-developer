package review

import (
	"strings"
	"testing"
)

func TestReport(t *testing.T) {
	report := Report()
	for _, want := range []string{
		"Capstone Final Review",
		"Go Service Capstone",
		"Cache layer",
		"caching-messaging",
		"Hardening categories:",
		"Concepts tied:",
	} {
		if !strings.Contains(report, want) {
			t.Fatalf("report missing %q\n%s", want, report)
		}
	}
}
