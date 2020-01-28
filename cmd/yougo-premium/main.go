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
	subManager, ytManager, downloader := initDependencies()
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
					chanName, err := commands.AddSubscription(c.Args().First(), subManager, ytManager)
					if err != nil {
						return err
					}
					err = subManager.SaveSubscriptions()
					if err != nil {
						return err
					}
					fmt.Println("Subscribed to " + chanName)
					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r"},
				Usage:   "remove a channel to your subscriptions via channel ID",
				Action: func(c *cli.Context) error {
					chanName, err := commands.RemoveSubscription(c.Args().First(), subManager)
					if err != nil {
						return err
					}
					err = subManager.SaveSubscriptions()
					if err != nil {
						return err
					}
					fmt.Println("Unsubcribed from " + chanName)
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list current subscriptions",
				Action: func(c *cli.Context) error {
					subs := commands.ListSubscriptions(subManager)
					if len(subs) == 0 {
						fmt.Println("Not subscribed to any channels.")
					} else {
						fmt.Println(subs)
					}
					return nil
				},
			},
			{
				Name:  "refresh",
				Usage: "download latest videos from your subscriptions",
				Action: func(c *cli.Context) error {
					downloader.InitVideoDirectory()

					fmt.Println("Checking for new videos...")
					vidCount, err := commands.RefreshVideos(subManager, ytManager, downloader)
					if err != nil {
						return err
					}
					err = subManager.SaveSubscriptions()
					if err != nil {
						return err
					}
					fmt.Printf("%v new videos downloaded.", vidCount)
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

func initDependencies() (*subscriptions.SubManager, *youtube.Wrapper, *video.Downloader) {
	subManager, err := subscriptions.InitializeSubManager()
	if err != nil {
		log.Fatal(err)
	}
	ytManager, err := youtube.InitializeWrapper()
	if err != nil {
		log.Fatal(err)
	}
	downloader, err := video.InitializeDownloader()
	if err != nil {
		log.Fatal(err)
	}
	return subManager, ytManager, downloader
}
