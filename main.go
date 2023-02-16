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
		Usage:       "keyfile filepath [--read] [--account]",
		Description: "Keychain-based file encryption",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "read",
				Usage:   "read a file",
				Aliases: []string{"r"},
			},
			&cli.StringFlag{
				Name:    "account",
				Usage:   "account name",
				Aliases: []string{"a"},
				EnvVars: []string{"KEYFILE_ACCOUNT"},
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return cli.ShowAppHelp(c)
			}
			path := c.Args().Get(0)
			read := c.Bool("read")
			account := c.String("account")

			var bs []byte
			var err error
			if read {
				bs, err = internal.DecodeFile(path, account)
			} else {
				bs, err = internal.EncodeFile(path, account)
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
