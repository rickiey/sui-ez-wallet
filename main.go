package main

import (
	"crypto/ed25519"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/rickiey/sui-ez-wallet/wallet"

	"github.com/urfave/cli/v2"
)

func main() {

	var suffix string
	var prefix string

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "gen",
				Aliases: []string{"g"},
				Usage:   "gen [flag]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "suffix",
						Usage: "suffix foo",
						// Value: suffix,
					},
					&cli.StringFlag{
						Name:  "prefix",
						Usage: "prefix boo",
						// Value: prefix,
					},
				},
				Action: func(cCtx *cli.Context) error {

					if suf := cCtx.String("suffix"); suf != "" {
						suffix = suf
					}

					if pre := cCtx.String("prefix"); pre != "" {
						prefix = pre
					}
					addr, privateKey := wallet.Gen(prefix, suffix)
					if addr == "" {
						fmt.Println("Your request is quite challenging; please try again, or change the prefix or suffix.")
						return nil
					}
					fmt.Printf("wallet address: %v\nprivate key: %v\n", addr, privateKey)
					return nil
				},
			},
			{
				Name:    "create-word",
				Aliases: []string{"c"},
				Usage:   "create-word <key-words>",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("create task: ")
					kword := cCtx.Args().First()
					hk := sha512.Sum512([]byte(kword))
					seed := hex.EncodeToString(hk[:32])
					privateKey, err := wallet.GenerateKeyFrmSeed(seed)
					if err != nil {
						panic(err)
					}
					publicKey := privateKey[32:]
					publicAddress := wallet.NewAddressByPublicKey(ed25519.PublicKey(publicKey))

					exprivate, err := wallet.ExportPrivatekey(privateKey)
					if err != nil {
						panic(err)
					}
					fmt.Printf("wallet address: %v\nprivate key: %v\n", publicAddress, exprivate)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
