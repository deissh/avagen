package commands

import "github.com/spf13/cobra"

type Command = cobra.Command

func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use:   "avatar",
	Short: "Open source avatar generator",
}
