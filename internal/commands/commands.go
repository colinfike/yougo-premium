package commands

import (
	"github.com/colinfike/yougo-premium/internal/subscriptions"
	"github.com/colinfike/yougo-premium/internal/youtube"
)

// TODO: Rename to app code?

// AddSubscription adds the channel associated with the URL to subscriptions
func AddSubscription(url string, subManager *subscriptions.SubManager, ytManager *youtube.YoutubeManager) error {
	channelInfo, err := ytManager.GetChannelInfo(url)
	if err != nil {
		return err
	}
	return subManager.AddSubscription(channelInfo)
}
func RemoveSubscription(id string, subManager *subscriptions.SubManager) error {
	return subManager.RemoveSubscription(id)
}
func ListSubscriptions(subManager *subscriptions.SubManager) string {
	return formatSubs(subManager.Subscriptions)
}
func ListVideos(subManager *subscriptions.SubManager) error {
	return nil
}
func RefreshVideos(subManager *subscriptions.SubManager, youtubeManger *youtube.YoutubeManager) error {
	return nil
}

func formatSubs(subMap map[string]subscriptions.Subscription) string {
	var output string
	for _, chanInfo := range subMap {
		output += chanInfo.ChannelName + "\n"
	}
	return output
}
