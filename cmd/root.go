package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tmplinux",
	Short: "tmplinux creates temporary virtual linux environments",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hi")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
