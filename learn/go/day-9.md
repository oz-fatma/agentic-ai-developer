# Day 9: Structs, Methods & Interfaces — Composition and Embedding

**Project:** Study Buddy — users, students, and admins via composition

## 1. Embed Structs

Go uses **embedding** instead of inheritance. Anonymous fields promote inner fields and methods:

```go
type User struct {
	ID    string
	Name  string
	Email string
}

type Student struct {
	User              // embedded — fields promoted
	GradeLevel int
	Major      string
}

type Admin struct {
	User
	Role string
}
```

Access promoted fields directly:

```go
student := Student{
	User:       User{ID: "u1", Name: "Alex", Email: "alex@school.edu"},
	GradeLevel: 3,
	Major:      "Computer Science",
}

fmt.Println(student.Name)  // "Alex" — promoted from User
fmt.Println(student.Email) // promoted
```

## 2. Compose Behavior

Build larger types from smaller ones:

```go
type StudyProfile struct {
	Student
	DailyGoalMinutes int
	TotalSessions    int
	PreferredSubject string
}

func (p *StudyProfile) ProgressReport() string {
	return fmt.Sprintf("%s (%s): %d sessions, goal %d min/day",
		p.Name, p.Major, p.TotalSessions, p.DailyGoalMinutes)
}
```

No base class — the relationship is explicit struct inclusion.

## 3. Override Methods

Outer struct methods **shadow** embedded methods with the same name:

```go
func (u User) Greet() string {
	return fmt.Sprintf("Hello, I'm %s", u.Name)
}

func (s Student) Greet() string {
	return fmt.Sprintf("Hi! I'm %s, year %d %s major",
		s.Name, s.GradeLevel, s.Major)
}

func (a Admin) Greet() string {
	return fmt.Sprintf("Admin %s (%s role)", a.Name, a.Role)
}
```

```go
student.Greet() // Student's version — embedded User.Greet is shadowed
```

Explicit call to embedded method:

```go
student.User.Greet() // "Hello, I'm Alex"
```

## 4. Design a Small Domain

Study Buddy user hierarchy:

```
User
├── ID, Name, Email
├── Greet()
│
├── Student (embeds User)
│   ├── GradeLevel, Major
│   └── Greet() — overridden
│
└── Admin (embeds User)
    ├── Role
    ├── Greet() — overridden
    └── CanManageCourses() bool
```

Full model:

```go
type Admin struct {
	User
	Role string
}

func (a Admin) CanManageCourses() bool {
	return a.Role == "superadmin" || a.Role == "course-admin"
}

func describeUser(u User) {
	fmt.Println(u.Greet())
}

func main() {
	student := Student{
		User: User{Name: "Alex"}, GradeLevel: 2, Major: "Go",
	}
	admin := Admin{
		User: User{Name: "Jordan"}, Role: "course-admin",
	}

	describeUser(student.User) // accepts User — composition enables this
	fmt.Println(admin.CanManageCourses()) // true
}
```

## Composition vs Inheritance

| Inheritance (other langs) | Go composition |
|---|---|
| Implicit "is-a" hierarchy | Explicit struct embedding |
| Deep override chains | Shadowing one level at a time |
| Fragile base classes | Small, independent structs |

## Summary

Study Buddy models `Student` and `Admin` by embedding `User` — reusing fields and methods without classical inheritance. Composition keeps relationships explicit and simple.
