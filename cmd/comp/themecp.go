package main

import (
	"github.com/spf13/cobra"
	"github.com/tomytp/icpc-companion/internal/runner"
)

var themecpCmd = &cobra.Command{
	Use:   "themecp [name]",
	Short: "Fetch problems from different contests into one folder (a, b, c, d)",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := ""
		if len(args) > 0 {
			name = args[0]
		}
		n, _ := cmd.Flags().GetInt("count")
		return runner.ThemeCP(name, n)
	},
}

func init() {
	themecpCmd.Flags().IntP("count", "n", 4, "Number of problems to fetch")
	rootCmd.AddCommand(themecpCmd)
}
