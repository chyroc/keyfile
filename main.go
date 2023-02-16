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
		Description: "Keychain-based file encryption",
		Commands: []*cli.Command{
			{
				Name:    "decrypt",
				Aliases: []string{"dec"},
				Usage:   "decrypt content from file",
				Flags: []cli.Flag{
					accountFlag,
					filepathFlag,
				},
				Action: func(c *cli.Context) error {
					account := c.String("account")
					path := c.String("file")

					bs, err := internal.DecryptFile(path, account)
					if err != nil {
						return err
					}
					fmt.Println(string(bs))
					fmt.Fprintf(os.Stderr, "decrypt file '%s' success\n", path)
					return nil
				},
			},
			{
				Name:    "encrypt",
				Aliases: []string{"enc"},
				Usage:   "encrypt content from file",
				Flags: []cli.Flag{
					accountFlag,
					filepathFlag,
				},
				Action: func(c *cli.Context) error {
					account := c.String("account")
					path := c.String("file")

					bs, err := internal.EncryptFile(path, account)
					if err != nil {
						return err
					}
					fmt.Println(string(bs))
					fmt.Fprintf(os.Stderr, "encrypt file '%s' success\n", path)
					return nil
				},
			},
			{
				Name:  "get-secret",
				Usage: "get secret from keychain",
				Flags: []cli.Flag{
					accountFlag,
				},
				Action: func(c *cli.Context) error {
					account := c.String("account")

					bs, err := internal.GetKeyChain(account)
					if err != nil {
						return err
					}
					fmt.Fprintf(os.Stderr, "get secret of '%s': '%s' success\n", account, string(bs))
					return nil
				},
			},
			{
				Name:  "set-secret",
				Usage: "set secret to keychain",
				Flags: []cli.Flag{
					accountFlag,
					secretFlag,
				},
				Action: func(c *cli.Context) error {
					account := c.String("account")
					secret := c.String("secret")

					err := internal.SetKeyChain(account, []byte(secret))
					if err != nil {
						return err
					}
					fmt.Fprintf(os.Stderr, "set secret of '%s': '%s' success\n", account, secret)
					return nil
				},
			},
			{
				Name:  "del-secret",
				Usage: "del secret from keychain",
				Flags: []cli.Flag{
					accountFlag,
				},
				Action: func(c *cli.Context) error {
					account := c.String("account")

					err := internal.DeleteKeyChain(account)
					if err != nil {
						return err
					}
					fmt.Fprintf(os.Stderr, "delete secret of '%s' success\n", account)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

var accountFlag = &cli.StringFlag{
	Name:     "account",
	Usage:    "account name",
	Aliases:  []string{"a"},
	EnvVars:  []string{"KEYFILE_ACCOUNT"},
	Required: true,
}

var filepathFlag = &cli.StringFlag{
	Name:     "file",
	Usage:    "filepath",
	Aliases:  []string{"f"},
	EnvVars:  []string{"KEYFILE_FILEPATH"},
	Required: true,
}

var secretFlag = &cli.StringFlag{
	Name:     "secret",
	Usage:    "secret to encrypt/decrypt file",
	Aliases:  []string{"s"},
	EnvVars:  []string{"KEYFILE_SECRET"},
	Required: true,
}
