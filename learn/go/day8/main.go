package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }

func (d Dog) Speak() string { return d.Name + " says woof" }

type Cat struct{ Name string }

func (c Cat) Speak() string { return c.Name + " says meow" }

func greet(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	animals := []Speaker{
		Dog{Name: "Rex"},
		Cat{Name: "Luna"},
	}
	for _, a := range animals {
		greet(a)
	}

	var s fmt.Stringer
	fmt.Println("nil interface:", s == nil)

	var speaker Speaker
	var dog *Dog
	speaker = dog
	fmt.Println("interface with nil concrete:", speaker == nil)
}
