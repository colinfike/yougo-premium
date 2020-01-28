package youtube

import (
	"context"

	"github.com/colinfike/yougo-premium/internal/config"
	"google.golang.org/api/option"
	youtubesdk "google.golang.org/api/youtube/v3"
)

type client interface {
	listChannels(string) (ChannelInfo, error)
	listVideos(string) (ChannelInfo, error)
	searchVideos(string, string, int64) ([]string, error)
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
