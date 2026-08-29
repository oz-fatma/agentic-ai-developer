package review

import (
	"strings"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day100/internal/capstone"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day100/internal/hardening"
)

// Report summarizes how the curriculum concepts connect in the capstone.
func Report() string {
	plan := capstone.DefaultPlan()
	byCat := hardening.ByCategory()

	var b strings.Builder
	b.WriteString("Capstone Final Review\n")
	b.WriteString("Plan: ")
	b.WriteString(plan.Title)
	b.WriteString("\nMilestones:\n")
	for _, m := range plan.Milestones {
		b.WriteString("  - ")
		b.WriteString(m.Name)
		b.WriteString(" (")
		b.WriteString(string(m.Phase))
		b.WriteString(")\n")
	}
	b.WriteString("Hardening categories:\n")
	for cat, items := range byCat {
		b.WriteString("  ")
		b.WriteString(cat)
		b.WriteString(": ")
		b.WriteString(itoa(len(items)))
		b.WriteString(" items\n")
	}
	b.WriteString("Concepts tied: cache, queue, pprof, semver, k8s, CI\n")
	return b.String()
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var digits []byte
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	return string(digits)
}
