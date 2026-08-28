package app

import (
	"fmt"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day56/internal/domain"
)

type Greeter struct{}

func NewGreeter() *Greeter {
	return &Greeter{}
}

func (g *Greeter) Greet(u domain.User) string {
	return fmt.Sprintf("Hello %s (%s)", u.Name, u.Role)
}
