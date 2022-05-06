package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
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
					id := uuid.New()
					fmt.Printf("uuid: %s\n", id)
					return nil
				},
			},
			{
				Name:  "ulid",
				Usage: "generate ulid",
				Action: func(c *cli.Context) error {
					t := time.Now()
					entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
					id := ulid.MustNew(ulid.Timestamp(t), entropy)
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
