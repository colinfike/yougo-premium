package subscriptions

import (
	"sync"
	"time"

	"github.com/colinfike/yougo-premium/internal/config"
)

type Subscription struct {
	ChannelID   string
	ChannelName string
	LastRefresh string
}

// Note: I think this level of DI may be overkill but I want to see how testing pans out compared to my non-DI golang work
type SubManager struct {
	config        config.Config
	Subscriptions map[string]Subscription
}

var once sync.Once
var manager *SubManager

func loadSubscriptions() map[string]Subscription {
	// Load in JSON subscriptions
	// Parse into a nice map and return
	return map[string]Subscription{}
}

func saveSubscriptions() {
	// iterate over map and encode every subscription as JSON and save the entire data
}

func GetSubManager(c *config.Config) *SubManager {
	once.Do(func() {
		manager = &SubManager{config: *c, Subscriptions: loadSubscriptions()}
	})
	return manager
}
func (subManager *SubManager) AddSubscription(channelID string) error {
	return nil
}
func (subManager *SubManager) RemoveSubscription(channelID string) error {
	return nil
}
func (subManager *SubManager) UpdateLastRefresh(channelID string, ts time.Time) error {
	return nil
}
func (subManager *SubManager) GetSubscriptions() ([]Subscription, error) {
	return []Subscription{}, nil
}
