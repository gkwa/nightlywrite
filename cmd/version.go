package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/gkwa/nightlywrite/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of nightlywrite",
	Long:  `All software has versions. This is nightlywrite's`,
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo := version.GetBuildInfo()
		fmt.Println(buildInfo)
	},
}
