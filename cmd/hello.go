package cmd

import (
	"os"

	"github.com/gkwa/nightlywrite/core"
	"github.com/spf13/cobra"
)

type helloCmd struct {
	ud core.UnicodeDetector
}

func NewHelloCmd(ud core.UnicodeDetector) *cobra.Command {
	hc := &helloCmd{ud: ud}

	cmd := &cobra.Command{
		Use:   "hello",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
		Run:   hc.run,
	}

	return cmd
}

func (hc *helloCmd) run(cmd *cobra.Command, args []string) {
	err := hc.ud.IsUnicode(os.Stdin, os.Stdout)
	if err != nil {
		cmd.PrintErrf("Error: %v\n", err)
		os.Exit(1)
	}
}
