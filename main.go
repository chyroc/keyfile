package main

import (
	"bytes"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"

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
					quietFlag,
					editorFlag,
				},
				Action: func(c *cli.Context) error {
					account := c.String("account")
					path := c.String("file")
					quiet := c.Bool("quiet")
					editor := c.String("editor")
					if editor != "" {
						quiet = true
					}

					bs, err := internal.DecryptFile(path, account)
					if err != nil {
						return err
					}
					if editor != "" {
						tmpFile, err := internal.WriteTempFile(bs)
						if err != nil {
							return err
						}
						defer func() { _ = os.Remove(tmpFile) }()

						go func() {
							if err := internal.ExecCommand(editor, tmpFile); err != nil {
								fmt.Printf("exec editor '%s' failed: %s", editor, err)
								os.Exit(1)
							}
						}()

						changed, err := internal.WaitFileChanged(tmpFile)
						if err != nil {
							return err
						}
						if <-changed {
							encryptData, err := internal.EncryptFile(tmpFile, account)
							if err != nil {
								return err
							}
							if bytes.Equal(bs, encryptData) {
								fmt.Printf("file '%s' not changed, skip write\n", path)
								return nil
							}

							fmt.Printf("file '%s' changed, write to file\n", path)

							return os.WriteFile(path, encryptData, 0o600)
						}
						return nil
					}
					fmt.Println(string(bs))
					if !quiet {
						_, _ = fmt.Fprintf(os.Stderr, "decrypt file '%s' success\n", path)
					}
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
					quietFlag,
				},
				Action: func(c *cli.Context) error {
					account := c.String("account")
					path := c.String("file")
					quiet := c.Bool("quiet")

					bs, err := internal.EncryptFile(path, account)
					if err != nil {
						return err
					}
					fmt.Println(string(bs))
					if !quiet {
						fmt.Fprintf(os.Stderr, "encrypt file '%s' success\n", path)
					}
					return nil
				},
			},
			{
				Name:  "get-secret",
				Usage: "get secret from keychain",
				Flags: []cli.Flag{
					accountFlag,
					quietFlag,
				},
				Action: func(c *cli.Context) error {
					account := c.String("account")
					quiet := c.Bool("quiet")

					bs, err := internal.GetKeyChain(account)
					if err != nil {
						return err
					}
					if !quiet {
						fmt.Fprintf(os.Stderr, "get secret of '%s': '%s' success\n", account, string(bs))
					} else {
						fmt.Println(string(bs))
					}
					return nil
				},
			},
			{
				Name:  "set-secret",
				Usage: "set secret to keychain",
				Flags: []cli.Flag{
					accountFlag,
					secretFlag,
					quietFlag,
				},
				Action: func(c *cli.Context) error {
					account := c.String("account")
					secret := c.String("secret")
					quiet := c.Bool("quiet")

					err := internal.SetKeyChain(account, []byte(secret))
					if err != nil {
						return err
					}
					if !quiet {
						fmt.Fprintf(os.Stderr, "set secret of '%s': '%s' success\n", account, secret)
					}
					return nil
				},
			},
			{
				Name:  "del-secret",
				Usage: "del secret from keychain",
				Flags: []cli.Flag{
					accountFlag,
					quietFlag,
				},
				Action: func(c *cli.Context) error {
					account := c.String("account")
					quiet := c.Bool("quiet")

					err := internal.DeleteKeyChain(account)
					if err != nil {
						return err
					}
					if !quiet {
						fmt.Fprintf(os.Stderr, "delete secret of '%s' success\n", account)
					}
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

var quietFlag = &cli.BoolFlag{
	Name:     "quiet",
	Usage:    "quiet mode",
	Aliases:  []string{"q"},
	EnvVars:  []string{"KEYFILE_QUIET"},
	Required: false,
}

var editorFlag = &cli.StringFlag{
	Name:     "editor",
	Usage:    "edit file with editor",
	Aliases:  []string{"e"},
	EnvVars:  []string{"KEYFILE_EDITOR"},
	Required: false,
}
