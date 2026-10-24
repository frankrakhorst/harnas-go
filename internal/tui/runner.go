package tui

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// RunNonInteractive reads lines from r, echoes handling slash commands, writes to w.
func RunNonInteractive(r io.Reader, w io.Writer) error {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	fmt.Fprint(w, "> ")
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		if cmd, ok := ParseCommand(line); ok {
			if !IsKnownCommand(cmd.Name) {
				fmt.Fprintf(w, "unknown command: /%s\n", cmd.Name)
				fmt.Fprint(w, "> ")
				continue
			}
			switch cmd.Name {
			case "help":
				fmt.Fprintln(w, "commands: /help /config /exit /model")
				fmt.Fprint(w, "> ")
			case "exit":
				return nil
			default:
				fmt.Fprintf(w, "ran /%s %s\n", cmd.Name, cmd.Args)
				fmt.Fprint(w, "> ")
			}
			continue
		}
		fmt.Fprintf(w, "you: %s\n", line)
		fmt.Fprint(w, "> ")
	}
	return sc.Err()
}
