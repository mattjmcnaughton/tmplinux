package cmd

import (
	"github.com/spf13/cobra"

	vh "github.com/mattjmcnaughton/tmplinux/pkg/virtualhost"
)

var containerSshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Ssh into container tmp linux environment",
	Run: func(cmd *cobra.Command, args []string) {
		container := vh.NewContainerVirtualHost()
		container.Ssh()
	},
}

var vmSshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Ssh into vm tmp linux environment",
	Run: func(cmd *cobra.Command, args []string) {
		vm := vh.NewVMVirtualHost()
		vm.Ssh()
	},
}

func init() {
	containerCmd.AddCommand(containerSshCmd)
	vmCmd.AddCommand(vmSshCmd)
}
