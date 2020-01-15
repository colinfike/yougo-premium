package youtube

import (
	"time"

	"github.com/colinfike/yougo-premium/internal/config"
)

// TODO: Let's try using wire here.
type YoutubeManager struct {
	config config.Config
	// CLient from SDK here?
}

// Return a struct with channelID and channelName
func (m *YoutubeManager) getChannelFromVideoID(videoID string) (string, error) {
	return "", nil
}

// Return a struct with channelID and channelName
func (m *YoutubeManager) getChannelFromChannelID(channelID string) (string, error) {
	return "", nil
}

func (m *YoutubeManager) fetchNewVideos(channelID string, ts time.Time) ([]string, error) {
	return []string{}, nil
}
