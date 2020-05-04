package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	FLAG_COUNT       = "count"
	FLAG_ALIAS_COUNT = "c"
)

func main() {
	app := &cli.App{
		Name:  "Repeat",
		Usage: "repeat characters",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    FLAG_COUNT,
				Value:   1,
				Usage:   "repeat count",
				Aliases: []string{FLAG_ALIAS_COUNT},
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return nil
			}
			fmt.Println(repeat(c.Args().Get(0), c.Int(FLAG_COUNT)))
			return nil
		},
		HideHelpCommand: true,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
