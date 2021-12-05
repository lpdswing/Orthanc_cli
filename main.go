package main

import (
	"fmt"
	"log"
	"ocli/core"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		{
			Name:  "listStudy",
			Usage: "Show the list of the studies",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "serve", Aliases: []string{"s"}},
				&cli.BoolFlag{Name: "option", Aliases: []string{"o"}},
				&cli.StringFlag{Name: "message", Aliases: []string{"m"}},
			},
			Action: func(c *cli.Context) error {
				fmt.Println(c.String("message"))
				core.ListStudy()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
