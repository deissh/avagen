package commands

import (
	"github.com/deissh/avagen/plugins"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

var ServerCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"s"},
	Short:   "Serve avatars",
	RunE:    AvatarServeCmdF,
}

func init() {
	ServerCmd.Flags().String("addr", "0.0.0.0:8080", "TCP address to listen to")
	ServerCmd.Flags().Bool("compress", false, "Whether to enable transparent response compression")

	RootCmd.AddCommand(ServerCmd)
}

func AvatarServeCmdF(command *cobra.Command, args []string) error {
	h := requestHandler
	compress, err := command.Flags().GetBool("compress")
	if err != nil {
		return err
	}
	if compress {
		h = fasthttp.CompressHandler(h)
	}

	addr, err := command.Flags().GetString("addr")
	if err != nil {
		return err
	}

	return fasthttp.ListenAndServe(addr, h)
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	// default values
	args := plugins.ParsedArg{
		"plugin": "identicon",
		"type":   "png",
		"name":   "Avatar Generator",
	}

	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		args[string(key)] = string(value)
	})

	plugin, err := plugins.Get(args["plugin"])
	if err != nil {
		ctx.SetStatusCode(http.StatusNotFound)
		return
	}

	data, err := plugin.Generate(args)
	if err != nil {
		log.Fatalln(err)
		return
	}

	ctx.SetBody(data)
	ctx.Response.Header.Set("Content-Type", "image/"+args["type"])
}
