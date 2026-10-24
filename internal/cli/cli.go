package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/frankrakhorst/harnas-go/internal/tui"
	"github.com/frankrakhorst/harnas-go/pkg/config"
	"github.com/frankrakhorst/harnas-go/pkg/harness"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "harnas",
	Short: "harnas-go AI harness",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("harnas-go 0.1.0")
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the harness",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}
		fmt.Printf("starting harness (log_level=%s)\n", cfg.LogLevel)
		return nil
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show config",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}
		fmt.Printf("log_level=%s\n", cfg.LogLevel)
		return nil
	},
}

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Run interactive TUI (non-interactive fallback)",
	RunE: func(cmd *cobra.Command, args []string) error {
		runner := &harness.Runner{Model: func(ctx context.Context, prompt string) (string, error) {
			return "echo: " + prompt, nil
		}}
		return tui.RunFullScreen(os.Stdin, os.Stdout, runner)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd, runCmd, configCmd, tuiCmd)
}

// Run executes the CLI.
func Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: harnas <run|version|config>")
	}
	rootCmd.SetArgs(args)
	return rootCmd.Execute()
}
