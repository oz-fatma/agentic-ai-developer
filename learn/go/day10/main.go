package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct{ Radius float64 }

func (c Circle) Area() float64      { return 3.14159 * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * 3.14159 * c.Radius }

type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(message string) { fmt.Println("[LOG]", message) }

type NoopLogger struct{}

func (NoopLogger) Log(message string) {}

func printShape(s Shape) {
	fmt.Printf("area=%.2f perimeter=%.2f\n", s.Area(), s.Perimeter())
}

func main() {
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 5},
	}

	tests := []struct {
		shape    Shape
		wantArea float64
	}{
		{Circle{Radius: 3}, 28.27},
		{Rectangle{Width: 4, Height: 5}, 20},
	}

	for _, tt := range tests {
		got := tt.shape.Area()
		fmt.Printf("area got=%.2f want~=%.2f\n", got, tt.wantArea)
	}

	for _, s := range shapes {
		printShape(s)
	}

	var logger Logger = ConsoleLogger{}
	logger.Log("shapes computed")
}
