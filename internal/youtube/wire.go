//+build wireinject

package youtube

import (
	"github.com/colinfike/yougo-premium/internal/config"
	"github.com/google/wire"
)

func InitializeWrapper() (*Wrapper, error) {
	wire.Build(config.NewConfig, NewWrapper)
	return &Wrapper{}, nil
}
