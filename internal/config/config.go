package config

import (
	"os"
	"sync"
)

// Config contains configuration info from env vars and user preferences.
type Config struct {
	DownloadLocation     string
	SubscriptionLocation string
	YoutubeAPIKey        string
}

const (
	defaultDownload             string = "/yougo-premium/"
	defaultSubscriptionLocation string = "/.yougo-premium"
)

var once sync.Once
var config *Config

func NewConfig() *Config {
	once.Do(func() {
		config = &Config{defaultDownload, defaultSubscriptionLocation, os.Getenv("YOUTUBE_API_KEY")}
	})
	return config
}
