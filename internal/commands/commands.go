package commands

import (
	"github.com/colinfike/yougo-premium/internal/subscriptions"
	"github.com/colinfike/yougo-premium/internal/youtube"
)

// Rename to app code?

// AddSubscription adds the channel associated with the URL to subscriptions
func AddSubscription(url string, subManager *subscriptions.SubManager, ytManager *youtube.YoutubeManager) error {
	// Check here if it's a youtube video link or a youtube channel link
	// Call the appropraite function in youtube pkg
	// Pass the following to subManager
	// ChannelID   string
	// ChannelName string
	// error otherwise

	// If url is a channel URL
	// call getChannelFromChannelID
	// if url is a ch
	ytManager.GetChannelInfo("https://www.youtube.com/channel/UCeBMccz-PDZf6OB4aV6a3eA")
	ytManager.GetChannelInfo("https://youtu.be/DkAgDThYsXE")
	ytManager.GetChannelInfo("https://www.youtube.com/watch?v=DkAgDThYsXE")
	return subManager.AddSubscription(url, "")
}
func RemoveSubscription(url string, subManager *subscriptions.SubManager) error {
	return nil
}
func ListSubscriptions(subManager *subscriptions.SubManager) error {
	return nil
}
func ListVideos() error {
	return nil
}
func RefreshVideos(subManager *subscriptions.SubManager, youtubeManger *youtube.YoutubeManager) error {
	return nil
}
