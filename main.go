package main

import (
	"os"

	"github.com/rarimo/broadcaster-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
