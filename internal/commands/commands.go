package commands

import (
	"fmt"
	"strings"
	"sync"

	"github.com/colinfike/yougo-premium/internal/subscriptions"
	"github.com/colinfike/yougo-premium/internal/video"
	"github.com/colinfike/yougo-premium/internal/youtube"
)

// AddSubscription adds the channel associated with the URL to subscriptions
func AddSubscription(url string, subManager *subscriptions.SubManager, ytManager *youtube.Wrapper) (string, error) {
	channelInfo, err := ytManager.GetChannelInfo(url)
	if err != nil {
		return "", err
	}
	return channelInfo.Name, subManager.AddSubscription(channelInfo.ID, channelInfo.Name)
}

// RemoveSubscription removes the channel associated with the channelID passed in
func RemoveSubscription(id string, subManager *subscriptions.SubManager) (string, error) {
	chanName, err := subManager.RemoveSubscription(id)
	if err != nil {
		return "", err
	}
	return chanName, nil
}

// ListSubscriptions lists all currently subscribed channels.
func ListSubscriptions(subManager *subscriptions.SubManager) string {
	return formatSubs(subManager.GetSubscriptions())
}

// ListVideos returns all currently downloaded videos.
func ListVideos(subManager *subscriptions.SubManager) error {
	return nil
}

// RefreshVideos downloads all new videos from all subscriptions since the last time they were refreshed.
func RefreshVideos(subManager *subscriptions.SubManager, youtubeManger *youtube.Wrapper, downloader *video.Downloader) (int, error) {
	var wg sync.WaitGroup
	var vidCount int

	for _, sub := range subManager.GetSubscriptions() {
		videos, err := youtubeManger.FetchNewVideos(sub.ChannelID, sub.LastRefresh)
		if err != nil {
			return 0, err
		}
		vidCount += len(videos)
		for _, video := range videos {
			wg.Add(1)
			go func(video youtube.VideoInfo) {
				defer wg.Done()
				fmt.Println("Downloading " + video.Name)
				err := downloader.DownloadVideo(video.ID)
				if err != nil {
					fmt.Println("Error downloading video id " + video.ID)
					// TODO: Handle errors here. Channels maybe?
				}
			}(video)
		}
		wg.Wait()

		subManager.UpdateLastRefresh(sub.ChannelID)
	}

	return vidCount, nil
}

func formatSubs(subMap map[string]subscriptions.Subscription) string {
	var output string
	for _, chanInfo := range subMap {
		output += chanInfo.ChannelName + "\n"
	}
	return strings.TrimSpace(output)
}
