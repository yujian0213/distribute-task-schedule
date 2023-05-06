package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
