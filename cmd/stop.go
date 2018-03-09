package cmd

import (
	"github.com/spf13/cobra"

	vh "github.com/mattjmcnaughton/tmplinux/pkg/virtualhost"
)

var containerStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a containerized tmplinux environment",
	Run: func(cmd *cobra.Command, args []string) {
		container := vh.NewContainerVirtualHost()
		container.Stop()
	},
}

var vmStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a vm tmplinux environment",
	Run: func(cmd *cobra.Command, args []string) {
		vm := vh.NewVMVirtualHost()
		vm.Stop()
	},
}

func init() {
	containerCmd.AddCommand(containerStopCmd)
	vmCmd.AddCommand(vmStopCmd)
}
