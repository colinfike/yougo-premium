package youtube

import (
	"errors"
	"regexp"
)

const (
	channelRegex  string = `youtube.com\/channel\/(.*)`
	videoRegex    string = `(?:youtu.be\/|youtube.com\/watch\?v=)([^&]*)`
	maxVideoCount int64  = 5
)

// ChannelInfo contains pertinent information for a Channel.
type ChannelInfo struct {
	ID, Name string
}

// VideoInfo contains pertinent information for a Video.
type VideoInfo struct {
	ID, Name string
}

// Wrapper is the main struct for interacting with the youtube package.
type Wrapper struct {
	youtubeClient client
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
func (m *Wrapper) FetchNewVideos(channelID, ts string) ([]VideoInfo, error) {
	videos, err := m.youtubeClient.searchVideos(channelID, ts, maxVideoCount)
	if err != nil {
		return nil, err
	}
	return videos, nil
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
