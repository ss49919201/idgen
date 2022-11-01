package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	if err := run(app()); err != nil {
		log.Fatal(err)
	}
}

func run(app *cli.App) error {
	return app.Run(os.Args)
}

func app() *cli.App {
	return &cli.App{
		Name:  "idgen",
		Usage: "Generate ID.",
		Commands: []*cli.Command{
			cmdGenerateUUID(),
			cmdGenerateULID(),
		},
		HideHelp: true,
	}
}
