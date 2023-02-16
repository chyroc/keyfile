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
		Usage:       "keyfile [--read] [--account] filepath",
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
				bs, err = internal.DecryptFile(path, account)
			} else {
				bs, err = internal.EncryptFile(path, account)
			}

			if err != nil {
				return err
			}
			fmt.Println(string(bs))
			if read {
				fmt.Fprintf(os.Stderr, "decrypt file '%s' success\n", path)
			} else {
				fmt.Fprintf(os.Stderr, "encrypt file '%s' success\n", path)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
