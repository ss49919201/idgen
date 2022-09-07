package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func cmdGenerateUUID() *cli.Command {
	return &cli.Command{
		Name:  "uuid",
		Usage: "Generate UUID.",
		Action: func(c *cli.Context) error {
			id := generateUUID()
			fmt.Printf("UUID: %s\n", id)
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
			fmt.Printf("ULID: %s\n", id)
			return nil
		},
	}
}
