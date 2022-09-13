package main

import (
	"os"

	"gitlab.com/tokend/bridge/core/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
