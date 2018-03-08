package cmd

import (
	"github.com/spf13/cobra"
)

var vmCmd = &cobra.Command{
	Use: "vm",
}

func init() {
	RootCmd.AddCommand(vmCmd)
}
