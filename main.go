package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	// FlagCount flag literal "count"
	FlagCount = "count"
	// FlagAliasCount flag alias for "count"
	FlagAliasCount = "c"
)

func main() {
	app := &cli.App{
		Name:  "Repeat",
		Usage: "repeat characters",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    FlagCount,
				Value:   1,
				Usage:   "repeat count",
				Aliases: []string{FlagAliasCount},
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return nil
			}
			fmt.Println(repeat(c.Args().Get(0), c.Int(FlagCount)))
			return nil
		},
		HideHelpCommand: true,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
