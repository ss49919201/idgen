package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "idgen",
		Usage: "generate id",
		Commands: []*cli.Command{
			{
				Name:  "uuid",
				Usage: "generate uuid",
				Action: func(c *cli.Context) error {
					id := generateUUID()
					fmt.Printf("uuid: %s\n", id)
					return nil
				},
			},
			{
				Name:  "ulid",
				Usage: "generate ulid",
				Action: func(c *cli.Context) error {
					id := generateULID()
					fmt.Printf("ulid: %s\n", id)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
