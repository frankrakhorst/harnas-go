package tui

import (
	"context"
	"fmt"
	"io"
	"os"
	"syscall"
	"unsafe"

	"github.com/frankrakhorst/harnas-go/pkg/harness"
)

// RunFullScreen renders a simple full-screen OpenCode-like layout.
func RunFullScreen(r io.Reader, w io.Writer, runner *harness.Runner) error {
	fd := int(os.Stdin.Fd())
	restore, err := RawMode(fd)
	if err != nil {
		return err
	}
	defer restore()

	var out []string
	var input string
	redraw := func() {
		fmt.Fprint(w, "\x1b[2J\x1b[H")
		fmt.Fprintln(w, "╭─ harnas ─────────────────────────────╮")
		h := 18
		start := 0
		if len(out) > h {
			start = len(out) - h
		}
		for i := start; i < len(out); i++ {
			fmt.Fprintf(w, "│ %s\n", out[i])
		}
		for i := len(out) - start; i < h; i++ {
			fmt.Fprintln(w, "│")
		}
		fmt.Fprintln(w, "╰──────────────────────────────────────╯")
		fmt.Fprintf(w, "❯ %s", input)
	}
	redraw()

	buf := make([]byte, 1)
	for {
		n, err := syscall.Read(fd, buf)
		if err != nil || n == 0 {
			break
		}
		b := buf[0]
		switch b {
		case '\r', '\n':
			if input == "" {
				continue
			}
			out = append(out, "you: "+input)
			if runner != nil {
				res, err := runner.Run(context.Background(), input)
				if err != nil {
					out = append(out, "error: "+err.Error())
				} else {
					out = append(out, "assistant: "+res)
				}
			}
			input = ""
		case 127, 8: // backspace
			if len(input) > 0 {
				input = input[:len(input)-1]
			}
		case 3: // Ctrl+C
			return nil
		case 27: // escape
			// read possible sequence
		default:
			if b >= 32 {
				input += string(b)
			}
		}
		redraw()
	}
	_ = io.EOF
	_ = unsafe.Pointer(nil)
	return nil
}