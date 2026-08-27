package main

import "fmt"

func incrementByValue(n int) {
	n++
}

func incrementByPointer(n *int) {
	*n++
}

func main() {
	x := 10
	fmt.Println("before value receiver:", x)
	incrementByValue(x)
	fmt.Println("after value receiver:", x)

	incrementByPointer(&x)
	fmt.Println("after pointer receiver:", x)

	p := new(int)
	*p = 42
	fmt.Println("new(int):", *p)

	q := &struct{ Name string }{Name: "Go"}
	fmt.Println("struct pointer:", q.Name)

	var nilPtr *int
	if nilPtr == nil {
		fmt.Println("nil pointer checked safely")
	}
}
