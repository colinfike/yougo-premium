package youtube

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var getChannelInfoTests = []struct {
	in  string
	out ChannelInfo
	err bool
}{
	{"https://www.youtube.com/channel/UCzPIY9Z0kRrJStr-5tbTjEQ", ChannelInfo{Name: "Kripparian", ID: "UCzPIY9Z0kRrJStr-5tbTjEQ"}, false},
	{"https://www.youtube.com/watch?v=u-SOcPQEgKw", ChannelInfo{Name: "Kripparian1", ID: "UCzPIY9Z0kRrJStr-5tbTjEQ"}, false},
	{"https://www", ChannelInfo{}, true},
	{"https://?v=u-SOcPQEgKw", ChannelInfo{}, true},
}

func TestGetChannelInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockClient := NewMockclient(mockCtrl)
	mockClient.EXPECT().listChannels("UCzPIY9Z0kRrJStr-5tbTjEQ").Return(ChannelInfo{Name: "Kripparian", ID: "UCzPIY9Z0kRrJStr-5tbTjEQ"}, nil)
	mockClient.EXPECT().listVideos("u-SOcPQEgKw").Return(ChannelInfo{Name: "Kripparian1", ID: "UCzPIY9Z0kRrJStr-5tbTjEQ"}, nil)

	wrapper, _ := NewWrapper(mockClient)
	for _, tt := range getChannelInfoTests {
		chanInfo, err := wrapper.GetChannelInfo(tt.in)
		assert.Equal(t, chanInfo, tt.out)
		if tt.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}
