package cmd

import (
	"github.com/spf13/cobra"

	vh "github.com/mattjmcnaughton/tmplinux/pkg/virtualhost"
)

var containerSSHCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH into container tmp linux environment",
	Run: func(cmd *cobra.Command, args []string) {
		container := vh.NewContainerVirtualHost()
		container.SSH()
	},
}

var vmSSHCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH into vm tmp linux environment",
	Run: func(cmd *cobra.Command, args []string) {
		vm := vh.NewVMVirtualHost()
		vm.SSH()
	},
}

func init() {
	containerCmd.AddCommand(containerSSHCmd)
	vmCmd.AddCommand(vmSSHCmd)
}
