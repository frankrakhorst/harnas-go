package cli

import (
	"fmt"
)

// Run dispatches CLI subcommands. Human-interface only; delegates to core functions.
func Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: harnas <run|version|config>")
	}
	switch args[0] {
	case "version":
		fmt.Println("harnas-go 0.1.0")
		return nil
	case "config":
		fmt.Println("config: (not implemented)")
		return nil
	case "run":
		fmt.Println("run: (no backend configured)")
		return nil
	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}
