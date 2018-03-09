package cmd

import (
	"github.com/spf13/cobra"

	vh "github.com/mattjmcnaughton/tmplinux/pkg/virtualhost"
)

var containerRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove the container containing the tmplinux environment",
	Run: func(cmd *cobra.Command, args []string) {
		container := vh.NewContainerVirtualHost()
		container.Rm()
	},
}

var vmRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove the vm containing the tmp linux environment",
	Run: func(cmd *cobra.Command, args []string) {
		vm := vh.NewVMVirtualHost()
		vm.Rm()
	},
}

func init() {
	containerCmd.AddCommand(containerRmCmd)
	vmCmd.AddCommand(vmRmCmd)
}
