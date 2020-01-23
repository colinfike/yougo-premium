package subscriptions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSubscription(t *testing.T) {
	subManger := &SubManager{subscriptions: make(map[string]Subscription)}

	subManger.AddSubscription("channelIdentifier", "channelName")

	assert.Equal(t, subManger.subscriptions["channelIdentifier"], Subscription{ChannelID: "channelIdentifier", ChannelName: "channelName", LastRefresh: ""})
}
func TestRemoveSubscription(t *testing.T) {
	subMap := make(map[string]Subscription)
	subMap["channelIdentifier"] = Subscription{ChannelID: "123", ChannelName: "RemoveTest"}
	subManger := &SubManager{subscriptions: subMap}

	subManger.RemoveSubscription("channelIdentifier")

	assert.Equal(t, subManger.subscriptions["channelIdentifier"], Subscription{})
}

func TestUpdateLastRefresh(t *testing.T) {
	subMap := make(map[string]Subscription)
	subMap["channelIdentifier"] = Subscription{ChannelID: "123", ChannelName: "RemoveTest"}
	subManger := &SubManager{subscriptions: subMap}

	subManger.UpdateLastRefresh("channelIdentifier")

	assert.NotEqual(t, subManger.subscriptions["channelIdentifier"].LastRefresh, "")
}

func TestGetSubscriptions(t *testing.T) {
	subMap := make(map[string]Subscription)
	subMap["channelIdentifier"] = Subscription{ChannelID: "123", ChannelName: "RemoveTest"}
	subManger := &SubManager{subscriptions: subMap}

	subs := subManger.GetSubscriptions()

	assert.Equal(t, subs, subMap)
}
