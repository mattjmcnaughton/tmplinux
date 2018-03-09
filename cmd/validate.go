package cmd

import (
	"github.com/spf13/cobra"

	vh "github.com/mattjmcnaughton/tmplinux/pkg/virtualhost"
)

// @TODO(mattjmcnaughton) Decide if validate show be a private command I call
// whenever a use attempts a container/vm action.
var containerValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate have proper deps for using tmplinux with container",
	Run: func(cmd *cobra.Command, args []string) {
		container := vh.NewContainerVirtualHost()
		container.Validate()
	},
}

var vmValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate have proper deps for using tmp linux with vm",
	Run: func(cmd *cobra.Command, args []string) {
		vm := vh.NewVMVirtualHost()
		vm.Validate()
	},
}

func init() {
	containerCmd.AddCommand(containerValidateCmd)
	vmCmd.AddCommand(vmValidateCmd)
}
