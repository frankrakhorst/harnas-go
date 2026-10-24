package cli

import (
	"fmt"

	"github.com/frankrakhorst/harnas-go/pkg/config"
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

func init() {
	rootCmd.AddCommand(versionCmd, runCmd, configCmd)
}

// Run executes the CLI.
func Run(args []string) error {
	rootCmd.SetArgs(args)
	return rootCmd.Execute()
}
