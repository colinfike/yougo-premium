package subscriptions

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"sync"
	"time"

	"github.com/colinfike/yougo-premium/internal/config"
	"github.com/colinfike/yougo-premium/internal/youtube"
)

type Subscription struct {
	ChannelID   string `json:"ChannelID"`
	ChannelName string `json:"ChannelName"`
	LastRefresh string `json:"LastRefresh"`
}

// Note: I think this level of DI may be overkill but I want to see how testing pans out compared to my non-DI golang work
type SubManager struct {
	config        config.Config
	Subscriptions map[string]Subscription
}

var once sync.Once
var manager *SubManager

func (subManager *SubManager) loadSubscriptions() (map[string]Subscription, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	newData, err := ioutil.ReadFile(user.HomeDir + subManager.config.SubscriptionLocation)
	if err != nil {
		return map[string]Subscription{}, nil
	}
	subscriptions := make(map[string]Subscription)
	err = json.Unmarshal(newData, &subscriptions)
	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (subManager *SubManager) saveSubscriptions() error {
	b, err := json.Marshal(subManager.Subscriptions)
	if err != nil {
		return err
	}
	user, err := user.Current()
	if err != nil {
		return err
	}
	ioutil.WriteFile(user.HomeDir+subManager.config.SubscriptionLocation, b, os.FileMode(int(0777)))
	return nil
}

func NewSubManager(c *config.Config) (*SubManager, error) {
	var err error
	once.Do(func() {
		manager = &SubManager{config: *c}
		var subs map[string]Subscription
		subs, err = manager.loadSubscriptions()
		manager.Subscriptions = subs
	})
	if err != nil {
		return &SubManager{}, err
	}
	return manager, nil
}

func (subManager *SubManager) AddSubscription(channel youtube.ChannelInfo) error {
	if _, ok := subManager.Subscriptions[channel.ID]; ok {
		return errors.New("Already subscribed to this channel")
	}
	subManager.Subscriptions[channel.ID] = Subscription{channel.ID, channel.Name, ""}
	err := subManager.saveSubscriptions()
	if err != nil {
		return err
	}
	return nil
}

func (subManager *SubManager) RemoveSubscription(channelID string) error {
	delete(subManager.Subscriptions, channelID)
	err := subManager.saveSubscriptions()
	if err != nil {
		return err
	}
	return nil
}

func (subManager *SubManager) UpdateLastRefresh(channelID string, ts time.Time) error {
	return nil
}

func (subManager *SubManager) GetSubscriptions() ([]Subscription, error) {
	return []Subscription{}, nil
}
