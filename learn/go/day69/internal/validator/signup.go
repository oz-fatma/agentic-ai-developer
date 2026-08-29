package validator

import (
	"fmt"
	"strings"
)

type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

type ValidationErrors []FieldError

func (ve ValidationErrors) Error() string {
	msgs := make([]string, len(ve))
	for i, e := range ve {
		msgs[i] = e.Error()
	}
	return strings.Join(msgs, "; ")
}

type Signup struct {
	Email    string
	Password string
	Age      int
}

func ValidateSignup(s Signup) error {
	var errs ValidationErrors
	if !strings.Contains(s.Email, "@") {
		errs = append(errs, FieldError{Field: "email", Message: "must contain @"})
	}
	if len(s.Password) < 8 {
		errs = append(errs, FieldError{Field: "password", Message: "min 8 chars"})
	}
	if s.Age < 13 {
		errs = append(errs, FieldError{Field: "age", Message: "must be >= 13"})
	}
	if len(errs) == 0 {
		return nil
	}
	return errs
}
