package main

import (
    "fmt"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "comp",
    Short: "ICPC Companion CLI",
    Long:  "Competitive programming problem downloader and tester",
}

// Execute runs the root command.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}

