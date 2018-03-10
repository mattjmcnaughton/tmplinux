package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is the non-executable root command to which we attach all other
// subcommads.
var RootCmd = &cobra.Command{
	Use: "tmplinux",
}
