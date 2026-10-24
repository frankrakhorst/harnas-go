package tui

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunNonInteractive(t *testing.T) {
	in := "/help\nhello\n/exit\n"
	var out bytes.Buffer
	if err := RunNonInteractive(strings.NewReader(in), &out); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "commands:") {
		t.Fatalf("missing help: %s", out.String())
	}
	if !strings.Contains(out.String(), "you: hello") {
		t.Fatalf("missing echo: %s", out.String())
	}
}

func TestSessionRoundTrip(t *testing.T) {
	dir := t.TempDir()
	p := dir + "/s.json"
	s := Session{Messages: []Message{{Role: "user", Content: "hi"}}}
	if err := SaveSession(p, s); err != nil {
		t.Fatal(err)
	}
	got, err := LoadSession(p)
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Messages) != 1 {
		t.Fatalf("got %d", len(got.Messages))
	}
}
