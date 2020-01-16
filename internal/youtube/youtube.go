package youtube

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/colinfike/yougo-premium/internal/config"
	"google.golang.org/api/option"
	youtubesdk "google.golang.org/api/youtube/v3"
)

type ChannelInfo struct {
	ID, Name string
}

type YoutubeManager struct {
	config         *config.Config
	youtubeService *youtubesdk.Service
}

const (
	channelRegex string = `youtube.com\/channel\/(.*)`
	videoRegex   string = `(?:youtu.be\/|youtube.com\/watch\?v=)([^=]*)`
)

// NewYoutubeManager initializes and returns a YoutubeManager
func NewYoutubeManager(c *config.Config) (*YoutubeManager, error) {
	ctx := context.Background()
	youtubeService, err := youtubesdk.NewService(ctx, option.WithAPIKey(c.YoutubeAPIKey))
	if err != nil {
		return nil, err
	}
	return &YoutubeManager{c, youtubeService}, nil
}

// GetChannelInfo accepts a video or channel URL and returns ChannelInfo
func (m *YoutubeManager) GetChannelInfo(url string) (ChannelInfo, error) {
	channelID, err := getMatch(channelRegex, url)
	if err == nil {
		return m.getChannelFromChannelID(channelID)
	}

	videoID, err := getMatch(videoRegex, url)
	if err == nil {
		return m.getChannelFromVideoID(videoID)
	}

	return ChannelInfo{}, errors.New("Invalid Channel/Video URL")
}

// Return a struct with channelID and channelName
func (m *YoutubeManager) getChannelFromVideoID(videoID string) (ChannelInfo, error) {
	return ChannelInfo{}, nil
}

// Return a struct with channelID and channelName
func (m *YoutubeManager) getChannelFromChannelID(channelID string) (ChannelInfo, error) {
	return ChannelInfo{}, nil
}

func (m *YoutubeManager) fetchNewVideos(channelID string, ts time.Time) ([]string, error) {
	return []string{}, nil
}

func getMatch(regex string, input string) (string, error) {
	re := regexp.MustCompile(regex)
	matches := re.FindSubmatch([]byte(input))
	if len(matches) > 0 {
		return string(matches[1]), nil
	}
	return "", errors.New("getMatch: No Match")
}
