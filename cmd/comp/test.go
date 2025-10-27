package main

import (
    "github.com/spf13/cobra"
    "github.com/tomytp/icpc-companion/internal/runner"
)

var (
    debugFlag bool
)

var testCmd = &cobra.Command{
    Use:   "test",
    Short: "Build latest .cpp and run tests",
    RunE: func(cmd *cobra.Command, args []string) error {
        return runner.Test(debugFlag)
    },
}

func init() {
    testCmd.Flags().BoolVarP(&debugFlag, "debug", "d", false, "enable debug build")
    rootCmd.AddCommand(testCmd)
}
