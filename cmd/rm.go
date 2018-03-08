package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var containerRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove the container containing the tmplinux environment",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("container rm called")
	},
}

var vmRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove the vm containing the tmp linux environment",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("vm rm called")
	},
}

func init() {
	containerCmd.AddCommand(containerRmCmd)
	vmCmd.AddCommand(vmRmCmd)
}
