package commands

import (
	"github.com/colinfike/yougo-premium/internal/subscriptions"
	"github.com/colinfike/yougo-premium/internal/youtube"
)

// Rename to app code?

// AddSubscription adds the channel associated with the URL to subscriptions
func AddSubscription(url string, subManager *subscriptions.SubManager, ytManager *youtube.YoutubeManager) error {
	channelInfo, err := ytManager.GetChannelInfo(url)
	if err != nil {
		return err
	}
	return subManager.AddSubscription(channelInfo)
}
func RemoveSubscription(url string, subManager *subscriptions.SubManager) error {
	return nil
}
func ListSubscriptions(subManager *subscriptions.SubManager) (map[string]subscriptions.Subscription, error) {
	return subManager.Subscriptions, nil
}
func ListVideos(subManager *subscriptions.SubManager) error {
	return nil
}
func RefreshVideos(subManager *subscriptions.SubManager, youtubeManger *youtube.YoutubeManager) error {
	return nil
}
