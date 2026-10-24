package tui

import (
	"fmt"
	"io"
	"os"
)

// RunAltScreen enters alt screen and runs non-interactive loop with prompt.
func RunAltScreen(r io.Reader, w io.Writer) error {
	fmt.Fprint(w, "\x1b[?1049h\x1b[H")
	fmt.Fprintln(w, "┌─ harnas ─────────────────────────────┐")
	fmt.Fprintln(w, "│  OpenCode-style TUI (look & feel)     │")
	fmt.Fprintln(w, "└──────────────────────────────────────┘")
	defer fmt.Fprint(w, "\x1b[?1049l")
	return RunNonInteractive(r, w)
}

// RunAltScreenDefault uses stdin/stdout.
func RunAltScreenDefault() error {
	return RunAltScreen(os.Stdin, os.Stdout)
}
