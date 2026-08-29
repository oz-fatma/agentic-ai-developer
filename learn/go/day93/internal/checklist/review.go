package checklist

// Items is a code review checklist for Go services.
var Items = []string{
	"Does the change have tests for new behavior?",
	"Are errors wrapped with context (%w)?",
	"Are goroutines bounded and shut down cleanly?",
	"Is sensitive data kept out of logs?",
	"Does API documentation match handler behavior?",
	"Are database migrations backward compatible?",
	"Will this work under expected concurrency load?",
}
