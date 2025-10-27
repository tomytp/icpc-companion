package main

import (
    "github.com/spf13/cobra"
    "github.com/tomytp/icpc-companion/internal/runner"
)

var runDebug bool

var runCmd = &cobra.Command{
    Use:   "run",
    Short: "Compile latest .cpp and run interactively (stdin/stdout)",
    RunE: func(cmd *cobra.Command, args []string) error {
        return runner.RunInteractive(runDebug)
    },
}

func init() {
    runCmd.Flags().BoolVarP(&runDebug, "debug", "d", false, "enable debug build")
    rootCmd.AddCommand(runCmd)
}
