package youtube

import (
	"context"
	"errors"
	"regexp"

	"github.com/colinfike/yougo-premium/internal/config"
	"google.golang.org/api/option"
	youtubesdk "google.golang.org/api/youtube/v3"
)

type ChannelInfo struct {
	ID, Name string
}

// YoutubeManager is the main struct for interacting with the youtube package. Contains the google
// api client for interacting with Youtube. // TODO: I think naming this Manager may be OK.
type YoutubeManager struct {
	config         *config.Config
	youtubeService *youtubesdk.Service
}

const (
	channelRegex string = `youtube.com\/channel\/(.*)`
	videoRegex   string = `(?:youtu.be\/|youtube.com\/watch\?v=)([^&]*)`
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

func (m *YoutubeManager) getChannelFromVideoID(videoID string) (ChannelInfo, error) {
	resp, err := m.youtubeService.Videos.List("snippet").Id(videoID).Do()
	if err != nil {
		return ChannelInfo{}, err
	}
	if len(resp.Items) == 0 {
		return ChannelInfo{}, errors.New("Video not found")
	}
	return ChannelInfo{resp.Items[0].Snippet.ChannelId, resp.Items[0].Snippet.ChannelTitle}, nil
}

func (m *YoutubeManager) getChannelFromChannelID(channelID string) (ChannelInfo, error) {
	resp, err := m.youtubeService.Channels.List("snippet").Id(channelID).Do()
	if err != nil {
		return ChannelInfo{}, err
	}
	if len(resp.Items) == 0 {
		return ChannelInfo{}, errors.New("Channel not found")
	}
	return ChannelInfo{resp.Items[0].Id, resp.Items[0].Snippet.Title}, nil
}

func (m *YoutubeManager) FetchNewVideos(channelID, ts string) ([]string, error) {
	req := m.youtubeService.Search.List("snippet").ChannelId(channelID).Order("date").MaxResults(5)
	if ts != "" {
		req.PublishedAfter(ts)
	}
	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	urls := make([]string, 0)
	for _, video := range resp.Items {
		urls = append(urls, video.Id.VideoId)
	}
	return urls, nil
}

// Move to some sort of utils package?
func getMatch(regex string, input string) (string, error) {
	re := regexp.MustCompile(regex)
	matches := re.FindSubmatch([]byte(input))
	if len(matches) > 0 {
		return string(matches[1]), nil
	}
	return "", errors.New("getMatch: No Match")
}
