package commands

import (
	"github.com/deissh/avagen/plugins"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"

	// load plugins
	_ "github.com/deissh/avagen/plugins/identicon"
)

var GenerateCmd = &cobra.Command{
	Use:     "generate [name]",
	Aliases: []string{"g"},
	Short:   "Generation",
	Args:    cobra.ExactArgs(1),
	RunE:    AvatarGenerateCmdF,
}

func init() {
	GenerateCmd.Flags().String("type", "png", "Optional image type; default png")
	GenerateCmd.Flags().String("plugin", "identicon", "Optional image style; default identicon")

	RootCmd.AddCommand(GenerateCmd)
}

func AvatarGenerateCmdF(command *cobra.Command, args []string) error {
	name := args[0]

	imageType, err := command.Flags().GetString("type")
	if err != nil {
		return errors.Wrap(err, "failed reading image type")
	}

	pluginName, err := command.Flags().GetString("plugin")
	if err != nil {
		return errors.Wrap(err, "failed reading image type")
	}

	plugin, err := plugins.Get(pluginName)
	if err != nil {
		return err
	}

	data, err := plugin.Generate(plugins.ParsedArg{"name": name, "type": imageType})
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(data)
	if err != nil {
		return err
	}

	return nil
}
