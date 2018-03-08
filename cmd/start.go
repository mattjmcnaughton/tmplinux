package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	vh "github.com/mattjmcnaughton/tmplinux/pkg/virtualhost"
)

var containerStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a containerized tmplinux environment",
	Run: func(cmd *cobra.Command, args []string) {
		container := vh.NewContainerVirtualHost()
		container.Start()
	},
}

var vmStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a vm tmplinux environment",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("vm start called")
	},
}

func init() {
	containerCmd.AddCommand(containerStartCmd)
	vmCmd.AddCommand(vmStartCmd)
}
