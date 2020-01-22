package youtube

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetChannelInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockClient := NewMockclient(mockCtrl)
	mockClient.EXPECT().listChannels("UCzPIY9Z0kRrJStr-5tbTjEQ").Return(ChannelInfo{Name: "Kripparian", ID: "UCzPIY9Z0kRrJStr-5tbTjEQ"}, nil)
	mockClient.EXPECT().listVideos("u-SOcPQEgKw").Return(ChannelInfo{Name: "Kripparian1", ID: "UCzPIY9Z0kRrJStr-5tbTjEQ"}, nil)

	wrapper, _ := NewWrapper(mockClient)
	chanInfo, _ := wrapper.GetChannelInfo("https://www.youtube.com/channel/UCzPIY9Z0kRrJStr-5tbTjEQ")

	assert.Equal(t, chanInfo.Name, "Kripparian")
	assert.Equal(t, chanInfo.ID, "UCzPIY9Z0kRrJStr-5tbTjEQ")

	chanInfo, _ = wrapper.GetChannelInfo("https://www.youtube.com/watch?v=u-SOcPQEgKw")

	assert.Equal(t, chanInfo.Name, "Kripparian1")
	assert.Equal(t, chanInfo.ID, "UCzPIY9Z0kRrJStr-5tbTjEQ")

}
