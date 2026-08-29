package api

import "context"

// User represents an account in the system.
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

// UserService defines read/update operations for users.
type UserService interface {
	// GetUser returns a user by ID or an error if not found.
	GetUser(ctx context.Context, id int) (User, error)
}

// Greeter says hello to a name.
//
// Example:
//
//	msg := Greeter("Go")
//	// msg == "Hello, Go!"
func Greeter(name string) string {
	return "Hello, " + name + "!"
}
