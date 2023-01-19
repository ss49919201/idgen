package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	app := &cli.App{
		Name:      "idgen",
		Usage:     "Generate ID.",
		UsageText: "idgen command",
		Commands: []*cli.Command{
			cmdGenerateUUID(),
			cmdGenerateULID(),
		},
		HideHelp: true,
	}
	return app.Run(os.Args)
}
