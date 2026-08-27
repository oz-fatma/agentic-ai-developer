package main

import "fmt"

type Course struct {
	Name     string
	Duration int
	Done     bool
}

func (c Course) Status() string {
	if c.Done {
		return "completed"
	}
	return "in progress"
}

func main() {
	c1 := Course{Name: "Go", Duration: 60}
	c2 := Course{Name: "Math", Duration: 45, Done: true}
	c3 := &Course{Name: "HTTP", Duration: 90}

	fmt.Println(c1)
	fmt.Println(c2.Status())
	fmt.Printf("%+v\n", c3)
	fmt.Println(c3.Status())
}
