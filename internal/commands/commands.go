package commands

import (
	"github.com/colinfike/yougo-premium/internal/subscriptions"
	"github.com/colinfike/yougo-premium/internal/youtube"
)

// Have main handle all the initialization and dependency injection
func addSubscription(url string, subManager *subscriptions.SubManager) error {
	return nil
}
func removeSubscription(url string, subManager *subscriptions.SubManager) error {
	return nil
}
func listSubscriptions(subManager *subscriptions.SubManager) error {
	return nil
}
func refreshVideos(subManager *subscriptions.SubManager, youtubeManger *youtube.YoutubeManager) error {
	return nil
}
