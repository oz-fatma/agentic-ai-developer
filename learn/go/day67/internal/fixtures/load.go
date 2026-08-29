package fixtures

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type UsersFile struct {
	Users []User `json:"users"`
}

func LoadUsersFromDir(dir string) ([]User, error) {
	data, err := os.ReadFile(filepath.Join(dir, "users.json"))
	if err != nil {
		return nil, err
	}
	var f UsersFile
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, err
	}
	return f.Users, nil
}

func LoadUsers(t *testing.T) []User {
	t.Helper()
	users, err := LoadUsersFromDir(testdataDir(t))
	if err != nil {
		t.Fatal(err)
	}
	return users
}

func testdataDir(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Join(filepath.Dir(file), "..", "testdata")
}
