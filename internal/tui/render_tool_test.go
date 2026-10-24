package tui

import (
	"bytes"
	"testing"
)

func TestRenderToolCall(t *testing.T) {
	var b bytes.Buffer
	RenderToolCall(&b, ToolCall{Name: "read", Args: "file.go"})
	if !bytes.Contains(b.Bytes(), []byte("[tool] read file.go")) {
		t.Fatalf("got %s", b.String())
	}
}
