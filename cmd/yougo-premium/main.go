package main

import (
	"fmt"
	"log"
	"os"

	"github.com/colinfike/yougo-premium/internal/commands"
	"github.com/colinfike/yougo-premium/internal/subscriptions"
	"github.com/colinfike/yougo-premium/internal/video"
	"github.com/colinfike/yougo-premium/internal/youtube"
	"github.com/urfave/cli/v2"
)

func main() {
	// TODO: Create function to initialize folder and file on first run if they don't exist
	// TODO: Overhaul errors in general
	// TODO: Error handling here
	subManager, _ := subscriptions.InitializeSubManager()
	// TODO: Error handling here
	ytManager, _ := youtube.InitializeYoutubeManager()
	// TODO: Error handling here
	downloader, _ := video.InitializeDownloader()
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
				Usage:   "remove a channel to your subscriptions via channel ID",
				Action: func(c *cli.Context) error {
					return commands.RemoveSubscription(c.Args().First(), subManager)
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list current subscriptions",
				Action: func(c *cli.Context) error {
					subs := commands.ListSubscriptions(subManager)
					fmt.Println(subs)
					return nil
				},
			},
			{
				Name:  "refresh",
				Usage: "download latest videos from your subscriptions",
				Action: func(c *cli.Context) error {
					return commands.RefreshVideos(subManager, ytManager, downloader)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
