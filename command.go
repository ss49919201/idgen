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

func cmdGenerateUUIDV7() *cli.Command {
	return &cli.Command{
		Name:  "uuidv7",
		Usage: "Generate UUID V7.",
		Action: func(c *cli.Context) error {
			id, err := generateUUIDV7()
			if err != nil {
				return err
			}
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
