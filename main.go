package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/chyroc/keyfile/internal"
)

func main() {
	app := &cli.App{
		Name:        "keyfile",
		Usage:       "keyfile filepath [--read]",
		Description: "Keychain-based file encryption",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "read",
				Usage:   "read a file",
				Aliases: []string{"r"},
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return cli.ShowAppHelp(c)
			}
			path := c.Args().Get(0)
			read := c.Bool("read")

			var bs []byte
			var err error
			if read {
				bs, err = internal.DecodeFile(path)
			} else {
				bs, err = internal.EncodeFile(path)
			}

			if err != nil {
				return err
			}
			fmt.Println(string(bs))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
