package main

import (
	"fmt"
	"os"

	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1-beta"
	app.CommandNotFound = func(ctx *cli.Context, cmd string) {
		fmt.Printf("No help topic for '%v'\n", cmd)
	}
	app.UseShortOptionHandling = true

	app.Commands = []*cli.Command{
		{
			Name:    "wallet",
			Aliases: []string{"w"},
			Usage:   "get wallet",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "mnemonic", Aliases: []string{"m"}, Required: false},
				&cli.StringFlag{Name: "seed", Aliases: []string{"s"}, Required: false},
			},
			Action: func(c *cli.Context) error {
				mnemonic := c.String("mnemonic")
				seed := c.String("seed")
				var w *wallet.Wallet
				var err error
				if seed == "" && mnemonic == "" {
					w, err = wallet.New()
				} else {
					w, err = wallet.Create(mnemonic, seed)
				}
				if err != nil {
					return err
				}

				fmt.Printf("%#v\n", w)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
