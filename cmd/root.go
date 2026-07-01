package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rice",
	Short: "RicePacker is a tool for sharing Linux desktop configurations",
	Long: `A tool that intelligently packs and installs dotfiles, handling path sanitization, 
hardware differences, and dependency management to easily share Linux desktop environments.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can be defined here, e.g., --verbose or --dry-run
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
}
