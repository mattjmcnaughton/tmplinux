package cmd

import (
	"github.com/spf13/cobra"
)

var containerCmd = &cobra.Command{
	Use: "container",
}

func init() {
	RootCmd.AddCommand(containerCmd)
}
