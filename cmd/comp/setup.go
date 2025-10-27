package main

import (
    "github.com/spf13/cobra"
    "github.com/tomytp/icpc-companion/internal/config"
    "github.com/tomytp/icpc-companion/internal/runner"
)

var setupCmd = &cobra.Command{
    Use:   "setup",
    Short: "Configure base path and templates",
    RunE: func(cmd *cobra.Command, args []string) error {
        return runner.Setup(config.DefaultPath())
    },
}

func init() {
    rootCmd.AddCommand(setupCmd)
}
