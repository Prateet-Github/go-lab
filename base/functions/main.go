package main

import (
	"fmt"
	"strings"
)

func FormatGreeting(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "Hello, world!"
	}

	runes := []rune(strings.ToLower(name))
	if len(runes) == 0 {
		return "Hello, world!"
	}

	runes[0] = []rune(strings.ToUpper(string(runes[0])))[0]
	return fmt.Sprintf("Hello, %s!", string(runes))
}
