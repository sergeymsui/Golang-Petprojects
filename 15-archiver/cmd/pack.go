package cmd

import "github.com/spf13/cobra"

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack programm",
}

func init() {
	rootCmd.AddCommand(packCmd)
}
