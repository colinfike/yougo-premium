//+build wireinject

package video

import (
	"github.com/colinfike/yougo-premium/internal/config"
	"github.com/google/wire"
)

func InitializeDownloader() (*Downloader, error) {
	wire.Build(config.NewConfig, NewDownloader)
	return &Downloader{}, nil
}
