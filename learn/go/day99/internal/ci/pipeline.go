package ci

// Step describes one stage in the CI pipeline summary.
type Step struct {
	Name    string
	Command string
}

// Pipeline returns the standard CI stages for the capstone repo.
func Pipeline() []Step {
	return []Step{
		{Name: "Checkout", Command: "actions/checkout@v4"},
		{Name: "Setup Go", Command: "actions/setup-go@v5"},
		{Name: "Format check", Command: "test -z \"$(gofmt -l .)\""},
		{Name: "Vet", Command: "go vet ./..."},
		{Name: "Test", Command: "go test ./..."},
		{Name: "Build", Command: "go build -o bin/api ./cmd/api"},
	}
}

func Summary() string {
	steps := Pipeline()
	return steps[0].Name + " -> " + steps[len(steps)-1].Name + " (" + itoa(len(steps)) + " stages)"
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
