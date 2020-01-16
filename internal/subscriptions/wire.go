//+build wireinject

package subscriptions

import (
	"github.com/colinfike/yougo-premium/internal/config"
	// "github.com/colinfike/yougo-premium/internal/subscriptions"
	"github.com/google/wire"
)

func InitializeSubManager() (*SubManager, error) {
	wire.Build(config.NewConfig, NewSubManager)
	return &SubManager{}, nil
}
