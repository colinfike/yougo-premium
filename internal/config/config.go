package config

import (
	"log"
	"os"
	"os/user"
	"sync"
)

// Config contains configuration info from env vars and user preferences.
type Config struct {
	DownloadLocation     string
	SubscriptionLocation string
	YoutubeAPIKey        string
	HomeDir              string
}

const (
	defaultDownload             string = "/yougo-premium/"
	defaultSubscriptionLocation string = "/.yougo-premium"
)

var once sync.Once
var config *Config

// NewConfig is the provider function for a Config object. Contains all configuration infomration used by other packages.
func NewConfig() *Config {
	once.Do(func() {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		config = &Config{defaultDownload, defaultSubscriptionLocation, os.Getenv("YOUTUBE_API_KEY"), usr.HomeDir}
	})
	return config
}
