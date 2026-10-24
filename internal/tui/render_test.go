package tui

import (
	"bytes"
	"testing"
)

func TestRenderHistory(t *testing.T) {
	var b bytes.Buffer
	RenderHistory(&b, Session{Messages: []Message{{Role: "user", Content: "hi"}}})
	if !bytes.Contains(b.Bytes(), []byte("[user] hi")) {
		t.Fatalf("got %s", b.String())
	}
}
