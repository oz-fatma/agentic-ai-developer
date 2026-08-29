package team

import (
	"fmt"
	"strings"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day95/internal/checklist"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day95/internal/version"
)

// ReleaseNotes builds a short release summary for reviewers.
func ReleaseNotes(feature string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Release %s\n", version.Current)
	fmt.Fprintf(&b, "Feature: %s\n", feature)
	fmt.Fprintln(&b, "Review checklist:")
	for i, item := range checklist.Items {
		fmt.Fprintf(&b, "  %d. %s\n", i+1, item)
	}
	return b.String()
}
