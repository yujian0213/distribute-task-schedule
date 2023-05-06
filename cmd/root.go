package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "crocodile",
	Short: "crocodile",
	Long:  "crocodile",
}

func Execute() {
	rootCmd.Execute()
}
