package main

import (
    "github.com/spf13/cobra"
    "github.com/tomytp/icpc-companion/internal/runner"
)

var solveCmd = &cobra.Command{
    Use:   "solve",
    Short: "Listen for problems and create folders",
    RunE: func(cmd *cobra.Command, args []string) error {
        return runner.Solve()
    },
}

func init() {
    rootCmd.AddCommand(solveCmd)
}
