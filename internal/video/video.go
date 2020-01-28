package video

import (
	"errors"
	"io"
	"net/http"
	"os"

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
	// TODO: Add configuration of download quality
	// bestFormat := vid.Formats.Best(ytdl.FormatResolutionKey)[0]
	url, err := vid.GetDownloadURL(vid.Formats[0])
	if err != nil {
		return "", err
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	name, path, err := dl.generateNameAndPath(vid)
	if err != nil {
		return "", err
	}
	destFile, err := os.Create(path)
	defer destFile.Close()
	if err != nil {
		return "", err
	}

	err = streamDownload(resp.Body, destFile)

	return name, nil
}

func (dl *Downloader) generateNameAndPath(vid *ytdl.VideoInfo) (string, string, error) {
	name := vid.Uploader + " - " + vid.Title + "." + vid.Formats[0].Extension
	path := dl.config.HomeDir + dl.config.DownloadLocation + name

	return name, path, nil
}

func streamDownload(body io.ReadCloser, file *os.File) error {
	buf := make([]byte, 102400)
	for {
		n, err := body.Read(buf)
		if err == io.EOF {
			file.Write(buf[0:n])
			break
		} else if err != nil {
			return err
		}
		file.Write(buf[0:n])
	}
	return nil
}

// InitVideoDirectory initializes the download directory if it does not exist
func (dl *Downloader) InitVideoDirectory() error {
	if _, err := os.Stat(dl.config.HomeDir + dl.config.DownloadLocation); os.IsNotExist(err) {
		os.Mkdir(dl.config.HomeDir+dl.config.DownloadLocation, os.FileMode(int(0755)))
	} else {
		return err
	}
	return nil
}
