package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var containerValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate have proper deps for using tmplinux with container",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("containerized validate called")
	},
}

var vmValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate have proper deps for using tmp linux with vm",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("vm validate called")
	},
}

func init() {
	containerCmd.AddCommand(containerValidateCmd)
	vmCmd.AddCommand(vmValidateCmd)
}
