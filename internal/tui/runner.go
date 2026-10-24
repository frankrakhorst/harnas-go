package tui

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/frankrakhorst/harnas-go/pkg/harness"
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

// RunNonInteractiveWithHarness runs the loop using a harness runner for chat turns.
func RunNonInteractiveWithHarness(ctx context.Context, r io.Reader, w io.Writer, runner *harness.Runner) error {
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
			case "exit":
				return nil
			default:
				fmt.Fprintf(w, "ran /%s %s\n", cmd.Name, cmd.Args)
			}
			fmt.Fprint(w, "> ")
			continue
		}
		out, err := runner.Run(ctx, line)
		if err != nil {
			fmt.Fprintf(w, "error: %v\n", err)
		} else {
			fmt.Fprintf(w, "assistant: %s\n", out)
		}
		fmt.Fprint(w, "> ")
	}
	return sc.Err()
}
