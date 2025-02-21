package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/urfave/cli/v2"
)

var writer = os.Stdout

var numberOfIDsFlag = &cli.IntFlag{
	Name:        "number",
	Aliases:     []string{"n"},
	Usage:       "Number of IDs to generate",
	Value:       1,
	DefaultText: "1",
}

func parseNumberOfIDsFlag(c *cli.Context) (*int, error) {
	n := c.String("number")

	parsedN, err := strconv.Atoi(n)
	if err != nil {
		return nil, fmt.Errorf("failed to parse number flag: %w", err)
	}

	if parsedN <= 0 {
		return nil, errors.New("number flag must be greater than 0")
	}

	return &parsedN, nil
}

func cmdGenerateUUID() *cli.Command {
	return &cli.Command{
		Name:  "uuid",
		Usage: "Generate UUID.",
		Flags: []cli.Flag{
			numberOfIDsFlag,
		},
		Action: func(c *cli.Context) error {
			number, err := parseNumberOfIDsFlag(c)
			if err != nil {
				return err
			}

			ids := make([]uuid.UUID, 0, *number)
			for i := 0; i < *number; i++ {
				ids = append(ids, generateUUID())
			}

			for i := 0; i < len(ids); i++ {
				fmt.Fprintln(writer, ids[i])
			}

			return nil
		},
	}
}

func cmdGenerateUUIDV7() *cli.Command {
	return &cli.Command{
		Name:  "uuidv7",
		Usage: "Generate UUID V7.",
		Flags: []cli.Flag{
			numberOfIDsFlag,
		},
		Action: func(c *cli.Context) error {
			number, err := parseNumberOfIDsFlag(c)
			if err != nil {
				return err
			}

			ids := make([]uuid.UUID, 0, *number)
			for i := 0; i < *number; i++ {
				id, err := generateUUIDV7()
				if err != nil {
					return err
				}
				ids = append(ids, id)
			}

			for i := 0; i < len(ids); i++ {
				fmt.Fprintln(writer, ids[i])
			}

			return nil
		},
	}
}

func cmdGenerateULID() *cli.Command {
	return &cli.Command{
		Name:  "ulid",
		Usage: "Generate ULID.",
		Flags: []cli.Flag{
			numberOfIDsFlag,
		},
		Action: func(c *cli.Context) error {
			number, err := parseNumberOfIDsFlag(c)
			if err != nil {
				return err
			}

			ids := make([]ulid.ULID, 0, *number)
			for i := 0; i < *number; i++ {
				ids = append(ids, generateULID())
			}

			for i := 0; i < len(ids); i++ {
				fmt.Fprintln(writer, ids[i])
			}

			return nil
		},
	}
}
