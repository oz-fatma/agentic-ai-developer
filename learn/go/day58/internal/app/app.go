package app

import "github.com/oz-fatma/agentic-ai-developer/learn/go/day58/internal/service"

type App struct {
	greeter service.Greeter
}

func NewApp(name string) *App {
	return &App{greeter: service.NewSimpleGreeter("Hello")}
}

func (a *App) Run() string {
	return a.greeter.Greet("DI")
}
