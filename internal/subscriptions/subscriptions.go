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

// Subscription contains all pertinent information about a subscribed channel
type Subscription struct {
	ChannelID   string `json:"ChannelID"`
	ChannelName string `json:"ChannelName"`
	LastRefresh string `json:"LastRefresh"`
}

// SubManager is the main interface with which you interfact with this package.
type SubManager struct {
	config        config.Config
	subscriptions map[string]Subscription
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
	b, err := json.Marshal(subManager.subscriptions)
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

// NewSubManager is the provider function for a SubManager.
func NewSubManager(c *config.Config) (*SubManager, error) {
	var err error
	once.Do(func() {
		manager = &SubManager{config: *c}
		var subs map[string]Subscription
		subs, err = manager.loadSubscriptions()
		manager.subscriptions = subs
	})
	if err != nil {
		return &SubManager{}, err
	}
	return manager, nil
}

// AddSubscription adds the new channel to subscriptions and writes them to disk.
func (subManager *SubManager) AddSubscription(channel youtube.ChannelInfo) error {
	if _, ok := subManager.subscriptions[channel.ID]; ok {
		return errors.New("Already subscribed to this channel")
	}
	subManager.subscriptions[channel.ID] = Subscription{channel.ID, channel.Name, ""}
	err := subManager.saveSubscriptions()
	if err != nil {
		return err
	}
	return nil
}

// RemoveSubscription removes the passed subscription and writes them to disk.
func (subManager *SubManager) RemoveSubscription(channelID string) (string, error) {
	chanName := subManager.subscriptions[channelID].ChannelName
	delete(subManager.subscriptions, channelID)
	err := subManager.saveSubscriptions()
	if err != nil {
		return "", err
	}
	return chanName, nil
}

// UpdateLastRefresh updates the timestamp of the last refresh for the subscription that matches the passed ChannelID
func (subManager *SubManager) UpdateLastRefresh(channelID string) error {
	sub := subManager.subscriptions[channelID]
	sub.LastRefresh = time.Now().Format(time.RFC3339)
	subManager.subscriptions[channelID] = sub

	err := subManager.saveSubscriptions()
	if err != nil {
		return err
	}

	return nil
}

// GetSubscriptions is a getter for the users current subscriptions.
func (subManager *SubManager) GetSubscriptions() map[string]Subscription {
	return subManager.subscriptions
}
