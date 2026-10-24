package tui

import (
	"fmt"
	"io"
)

// RenderHistory writes session messages to w.
func RenderHistory(w io.Writer, s Session) {
	for _, m := range s.Messages {
		fmt.Fprintf(w, "[%s] %s\n", m.Role, m.Content)
	}
}

// ToolCall represents a tool invocation.
type ToolCall struct {
	Name string
	Args string
}

// RenderToolCall writes a tool call to w.
func RenderToolCall(w io.Writer, tc ToolCall) {
	fmt.Fprintf(w, "[tool] %s %s\n", tc.Name, tc.Args)
}
