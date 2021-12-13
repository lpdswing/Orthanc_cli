package main

import (
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
			Name:    "delOverview",
			Aliases: []string{"do"},
			Usage:   "Delete biomind overview from orthanc.",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "study_uid", Aliases: []string{"s"}},
			},
			Action: func(c *cli.Context) error {
				studyUid := c.String("study_uid")
				if studyUid != "" {
					core.DeleteOverviewByStudy(studyUid)
					return nil
				}
				core.DeleteAllOverview()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
