package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func loadUser(id int) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("load user %d: %w", id, ErrNotFound)
	}
	return "Alex", nil
}

func main() {
	name, err := loadUser(1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("user:", name)

	_, err = loadUser(-1)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("handled:", err)
		}
	}
}
