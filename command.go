package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var writer = os.Stdout

func cmdGenerateUUID() *cli.Command {
	return &cli.Command{
		Name:  "uuid",
		Usage: "Generate UUID.",
		Action: func(c *cli.Context) error {
			id := generateUUID()
			fmt.Fprintln(writer, id)
			return nil
		},
	}
}

func cmdGenerateULID() *cli.Command {
	return &cli.Command{
		Name:  "ulid",
		Usage: "Generate ULID.",
		Action: func(c *cli.Context) error {
			id := generateULID()
			fmt.Fprintln(writer, id)
			return nil
		},
	}
}
