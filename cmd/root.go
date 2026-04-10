package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:"cmdock",
	Version:"1.0.6",
	Short:"cmdock is a command logger tool",
	Long:"cmdock is a command logger tool for tracking and managing command line history.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Oops. An error while executing cmdock'%s'\n", err)
        os.Exit(1)
    }
}