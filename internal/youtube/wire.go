//+build wireinject

package youtube

import (
	"github.com/colinfike/yougo-premium/internal/config"
	"github.com/google/wire"
)

func InitializeYoutubeManager() (*YoutubeManager, error) {
	wire.Build(config.NewConfig, NewYoutubeManager)
	return &YoutubeManager{}, nil
}
