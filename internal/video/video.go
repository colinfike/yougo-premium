package video

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/colinfike/yougo-premium/internal/config"
	"github.com/colinfike/ytdl"
)

// Downloader is the main interface into the video package.
type Downloader struct {
	config *config.Config
}

// NewDownloader is the provider function for a Downloader.
func NewDownloader(c *config.Config) (*Downloader, error) {
	return &Downloader{c}, nil
}

// DownloadVideo downloads a youtube video to the default download location.
func (dl *Downloader) DownloadVideo(videoID string) (string, error) {
	vid, err := ytdl.GetVideoInfo("https://www.youtube.com/watch?v=" + videoID)
	if err != nil {
		return "", errors.New("Failed to get video info")
	}

	// TODO: Can we stream this? I don't want to download into memory then write to disk
	// ToDo: I think we'll need to use a different libary to do the downloading from the URL
	buf := new(bytes.Buffer)
	// TODO: Without streaming this format is killing my laptop
	// bestFormat := vid.Formats.Best(ytdl.FormatResolutionKey)[0]
	err = vid.Download(vid.Formats[0], buf)
	if err != nil {
		return "", errors.New("Error downloading video")
	}

	user, err := user.Current()
	if err != nil {
		return "", err
	}

	name := vid.Uploader + " - " + vid.Title + "." + vid.Formats[0].Extension
	path := user.HomeDir + dl.config.DownloadLocation + name
	return name, ioutil.WriteFile(path, buf.Bytes(), os.FileMode(int(0777)))
}
