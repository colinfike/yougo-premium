package config

import "os"

// Config contains configuration info from env vars and user preferences.
type Config struct {
	DownloadLocation     string
	SubscriptionLocation string
	YoutubeAccessKey     string
	YoutubeSecretKey     string
}

const (
	defaultDownload             string = "/videos"
	defaultSubscriptionLocation string = "$home/.yougo-premium"
)

var config Config = Config{defaultDownload, defaultSubscriptionLocation, os.Getenv("YOUTUBE_ACCESS_KEY"), os.Getenv("YOUTUBE_SECRET_KEY")}

func getConfig() *Config {
	return &config
}
