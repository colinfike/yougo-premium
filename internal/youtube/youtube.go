package youtube

import (
	"context"
	"errors"
	"regexp"

	"github.com/colinfike/yougo-premium/internal/config"
	"google.golang.org/api/option"
	youtubesdk "google.golang.org/api/youtube/v3"
)

const (
	channelRegex  string = `youtube.com\/channel\/(.*)`
	videoRegex    string = `(?:youtu.be\/|youtube.com\/watch\?v=)([^&]*)`
	maxVideoCount int64  = 5
)

// ChannelInfo contains pertitnent information for a Channel.
type ChannelInfo struct {
	ID, Name string
}

// Wrapper is the main struct for interacting with the youtube package.
type Wrapper struct {
	youtubeClient client
}

type client interface {
	listChannels(string) (ChannelInfo, error)
	listVideos(string) (ChannelInfo, error)
	searchVideos(string, string, int64) ([]string, error)
}

// NewWrapper initializes and returns a Wrapper
func NewWrapper(client client) (*Wrapper, error) {
	return &Wrapper{client}, nil
}

// GetChannelInfo accepts a video or channel URL and returns ChannelInfo
func (m *Wrapper) GetChannelInfo(url string) (ChannelInfo, error) {
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

func (m *Wrapper) getChannelFromVideoID(videoID string) (ChannelInfo, error) {
	chanInfo, err := m.youtubeClient.listVideos(videoID)
	if err != nil {
		return ChannelInfo{}, err
	}
	if (ChannelInfo{}) == chanInfo {
		return ChannelInfo{}, errors.New("Channel associated with videoID " + videoID + " not found")
	}
	return chanInfo, nil
}

func (m *Wrapper) getChannelFromChannelID(channelID string) (ChannelInfo, error) {
	chanInfo, err := m.youtubeClient.listChannels(channelID)
	if err != nil {
		return ChannelInfo{}, err
	}
	if (ChannelInfo{}) == chanInfo {
		return ChannelInfo{}, errors.New("Channel associated with channelID " + channelID + " not found")
	}
	return chanInfo, nil
}

// FetchNewVideos returns an array of video URLs.
func (m *Wrapper) FetchNewVideos(channelID, ts string) ([]string, error) {
	urls, err := m.youtubeClient.searchVideos(channelID, ts, maxVideoCount)
	if err != nil {
		return nil, err
	}
	return urls, nil
}

// TODO: Move to some sort of utils package?
func getMatch(regex string, input string) (string, error) {
	re := regexp.MustCompile(regex)
	matches := re.FindSubmatch([]byte(input))
	if len(matches) > 0 {
		return string(matches[1]), nil
	}
	return "", errors.New("getMatch: No Match")
}

// youtubeClient is a wrapper around the youtube API in order to faciliate DI/unit testing of the functions on Wrapper.
// Wrapper barely does anything more than youtubeClient but the exercise is worth it.
type youtubeClient struct {
	service *youtubesdk.Service
	config  *config.Config
}

// NewYoutubeClient is the provider function for the wrapper around the Google Youtube API.
func NewYoutubeClient(c *config.Config) (*youtubeClient, error) {
	ctx := context.Background()
	youtubeService, err := youtubesdk.NewService(ctx, option.WithAPIKey(c.YoutubeAPIKey))
	if err != nil {
		return &youtubeClient{}, err
	}
	return &youtubeClient{youtubeService, c}, err
}

func (yt *youtubeClient) listChannels(channelID string) (ChannelInfo, error) {
	resp, err := yt.service.Channels.List("snippet").Id(channelID).Do()
	if err != nil || len(resp.Items) == 0 {
		return ChannelInfo{}, err
	}
	return ChannelInfo{resp.Items[0].Id, resp.Items[0].Snippet.Title}, nil
}

func (yt *youtubeClient) listVideos(videoID string) (ChannelInfo, error) {
	resp, err := yt.service.Videos.List("snippet").Id(videoID).Do()
	if err != nil || len(resp.Items) == 0 {
		return ChannelInfo{}, err
	}
	return ChannelInfo{resp.Items[0].Snippet.ChannelId, resp.Items[0].Snippet.ChannelTitle}, nil
}

func (yt *youtubeClient) searchVideos(channelID string, ts string, maxResults int64) ([]string, error) {
	req := yt.service.Search.List("snippet").ChannelId(channelID).Order("date").MaxResults(maxResults)
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
