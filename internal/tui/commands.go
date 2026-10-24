package tui

import "strings"

// Command represents a parsed slash command.
type Command struct {
	Name string
	Args string
}

// ParseCommand parses input. Returns ok=false if not a slash command.
func ParseCommand(input string) (Command, bool) {
	s := strings.TrimSpace(input)
	if !strings.HasPrefix(s, "/") {
		return Command{}, false
	}
	parts := strings.SplitN(strings.TrimPrefix(s, "/"), " ", 2)
	c := Command{Name: strings.ToLower(parts[0])}
	if len(parts) > 1 {
		c.Args = strings.TrimSpace(parts[1])
	}
	return c, true
}

// IsKnownCommand reports whether name is supported.
func IsKnownCommand(name string) bool {
	switch strings.ToLower(name) {
	case "help", "config", "exit", "model":
		return true
	}
	return false
}
