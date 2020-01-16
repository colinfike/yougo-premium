package main

import (
	"fmt"
	"log"
	"os"

	"github.com/colinfike/yougo-premium/internal/commands"
	"github.com/colinfike/yougo-premium/internal/subscriptions"
	"github.com/colinfike/yougo-premium/internal/youtube"
	"github.com/urfave/cli/v2"
)

func main() {
	// TODO: Error handling here
	subManager, _ := subscriptions.InitializeSubManager()
	// TODO: Error handling here
	ytManager, _ := youtube.InitializeYoutubeManager()
	app := &cli.App{
		Name:     "yougo-premium",
		Usage:    "Follow and download latest videos from subscriptions",
		HelpName: "yougo-premium",
		Commands: []*cli.Command{
			{
				Name:      "add",
				Aliases:   []string{"a"},
				Usage:     "add a channel to your subscriptions via channel url or video url",
				ArgsUsage: "[url]",
				Action: func(c *cli.Context) error {
					return commands.AddSubscription(c.Args().First(), subManager, ytManager)
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r"},
				Usage:   "remove a channel to your subscriptions",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list current subscriptions",
				Action: func(c *cli.Context) error {
					subs, err := commands.ListSubscriptions(subManager)
					if err != nil {
						return err
					}
					fmt.Println(subs)
					return nil
				},
			},
			{
				Name:  "refresh",
				Usage: "download latest videos from your subscriptions",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
