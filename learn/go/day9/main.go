package main

import "fmt"

type Person struct {
	Name  string
	Email string
}

type User struct {
	Person
	ID int
}

type Admin struct {
	User
	Role string
}

func (u User) Describe() string {
	return fmt.Sprintf("User %d: %s", u.ID, u.Name)
}

func (a Admin) Describe() string {
	return fmt.Sprintf("Admin %s (%s)", a.Name, a.Role)
}

func main() {
	user := User{
		Person: Person{Name: "Alex", Email: "alex@example.com"},
		ID:     1,
	}
	admin := Admin{
		User: User{
			Person: Person{Name: "Sam", Email: "sam@example.com"},
			ID:     2,
		},
		Role: "editor",
	}

	fmt.Println(user.Describe())
	fmt.Println(user.Email)
	fmt.Println(admin.Describe())
}
