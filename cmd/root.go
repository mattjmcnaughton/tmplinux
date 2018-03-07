package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mattjmcnaughton/tmplinux/pkg/virtualhost"
)

var rootCmd = &cobra.Command{
	Use:   "tmplinux",
	Short: "tmplinux creates temporary virtual linux environments",
	Run: func(cmd *cobra.Command, args []string) {
		vh := virtualhost.NewContainerVirtualHost()

		vh.Start()
		vh.Stop()
		vh.Ssh()
		vh.UniqueVirtualHostMethod()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
