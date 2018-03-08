package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var containerStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a containerized tmplinux environment",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("containerized stop called")
	},
}

var vmStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a vm tmplinux environment",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("vm stop called")
	},
}

func init() {
	containerCmd.AddCommand(containerStopCmd)
	vmCmd.AddCommand(vmStopCmd)
}
