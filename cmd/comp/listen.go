package main

import (
    "github.com/spf13/cobra"
    "github.com/tomytp/icpc-companion/internal/runner"
)

var listenCmd = &cobra.Command{
    Use:   "listen",
    Short: "Listen and print incoming HTTP payloads (dry mode)",
    RunE: func(cmd *cobra.Command, args []string) error {
        return runner.DryListen()
    },
}

func init() {
    rootCmd.AddCommand(listenCmd)
}
