package fixtures_test

import (
	"testing"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day67/internal/fixtures"
)

func TestLoadUsersFixture(t *testing.T) {
	users := fixtures.LoadUsers(t)
	if len(users) != 2 {
		t.Fatalf("want 2 users, got %d", len(users))
	}
	if users[0].Name != "Ada" {
		t.Fatalf("first user = %+v", users[0])
	}
}
