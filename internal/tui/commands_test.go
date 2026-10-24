package tui

import "testing"

func TestParseCommand(t *testing.T) {
	c, ok := ParseCommand("/help")
	if !ok || c.Name != "help" {
		t.Fatalf("got %+v %v", c, ok)
	}
	c, ok = ParseCommand("/model gpt-4")
	if !ok || c.Name != "model" || c.Args != "gpt-4" {
		t.Fatalf("got %+v %v", c, ok)
	}
	if _, ok := ParseCommand("hello"); ok {
		t.Fatal("expected false")
	}
}

func TestIsKnownCommand(t *testing.T) {
	for _, n := range []string{"help", "config", "exit", "model"} {
		if !IsKnownCommand(n) {
			t.Fatalf("expected known %s", n)
		}
	}
	if IsKnownCommand("nope") {
		t.Fatal("expected unknown")
	}
}
