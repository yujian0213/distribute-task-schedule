package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "run client",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client")
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
