package main

import (
	"github.com/deissh/avagen/commands"
	_ "github.com/deissh/avagen/plugins/identicon"
	"os"
)

func main() {
	if err := commands.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
