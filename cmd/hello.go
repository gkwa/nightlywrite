package cmd

import (
	"fmt"
	"os"

	"github.com/gkwa/nightlywrite/core"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		isUnicode, codePoint, err := core.IsUnicode("testdata/test.md")
		if err != nil {
			cmd.PrintErrf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Is Unicode: %v\nCode Point: %s\n", isUnicode, codePoint)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
