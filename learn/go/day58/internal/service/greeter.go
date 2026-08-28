package service

type Greeter interface {
	Greet(name string) string
}

type SimpleGreeter struct {
	prefix string
}

func NewSimpleGreeter(prefix string) *SimpleGreeter {
	return &SimpleGreeter{prefix: prefix}
}

func (g *SimpleGreeter) Greet(name string) string {
	return g.prefix + ", " + name + "!"
}
