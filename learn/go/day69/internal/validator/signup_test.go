package validator_test

import (
	"errors"
	"testing"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day69/internal/validator"
)

func TestValidateSignup(t *testing.T) {
	tests := []struct {
		name    string
		input   validator.Signup
		wantErr bool
	}{
		{
			name:    "valid",
			input:   validator.Signup{Email: "a@b.com", Password: "longenough", Age: 20},
			wantErr: false,
		},
		{
			name:    "bad email",
			input:   validator.Signup{Email: "bad", Password: "longenough", Age: 20},
			wantErr: true,
		},
		{
			name:    "short password",
			input:   validator.Signup{Email: "a@b.com", Password: "short", Age: 20},
			wantErr: true,
		},
		{
			name:    "young age",
			input:   validator.Signup{Email: "a@b.com", Password: "longenough", Age: 10},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := validator.ValidateSignup(tc.input)
			if tc.wantErr && err == nil {
				t.Fatal("expected error")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tc.wantErr {
				var ve validator.ValidationErrors
				if !errors.As(err, &ve) {
					t.Fatalf("want ValidationErrors, got %T", err)
				}
			}
		})
	}
}
