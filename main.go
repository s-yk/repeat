package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	// flagCount flag literal "count"
	flagCount = "count"
	// flagAliasCount flag alias for "count"
	flagAliasCount = "c"
)

func main() {
	app := &cli.App{
		Name:  "Repeat",
		Usage: "repeat characters",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    flagCount,
				Value:   1,
				Usage:   "repeat count",
				Aliases: []string{flagAliasCount},
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				scn := bufio.NewScanner(os.Stdin)
				for scn.Scan() {
					if err := scn.Err(); err != nil {
						return err
					}
					fmt.Println(repeat(scn.Text(), c.Int(flagCount)))
				}
				return nil
			}
			fmt.Println(repeat(c.Args().Get(0), c.Int(flagCount)))
			return nil
		},
		HideHelpCommand: true,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
