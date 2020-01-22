//+build wireinject

package youtube

import (
	"github.com/colinfike/yougo-premium/internal/config"
	"github.com/google/wire"
)

var NewYoutubeClientSet = wire.NewSet(
    NewYoutubeClient,
    wire.Bind(new(client), new(*youtubeClient)))

func InitializeWrapper() (*Wrapper, error) {
	wire.Build(config.NewConfig, NewYoutubeClientSet, NewWrapper)
	return &Wrapper{}, nil
}
func InitializeYoutubeClient() (*youtubeClient, error) {
	wire.Build(config.NewConfig, NewYoutubeClient)
	return &youtubeClient{}, nil
}

