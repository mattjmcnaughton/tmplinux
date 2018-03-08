package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var containerSshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Ssh into container tmp linux environment",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("container ssh called")
	},
}

var vmSshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Ssh into vm tmp linux environment",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("vm ssh called")
	},
}

func init() {
	containerCmd.AddCommand(containerSshCmd)
	vmCmd.AddCommand(vmSshCmd)
}
