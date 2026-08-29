package greet

import (
	"strings"
	"unicode"
)

func Title(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "Friend"
	}
	runes := []rune(strings.ToLower(name))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func Message(name string) string {
	return "Hello, " + Title(name) + "!"
}
