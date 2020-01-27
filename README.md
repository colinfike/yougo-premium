# Yougo-Premium

Yougo-Premium is a command line tool written in Go that allows you to subscribe to Youtube channels and download their latest videos to your machine. It tracks the last time you downloaded new videos for a channel so you only get the newest videos.

## Getting Started

Currently the tool is only configured to work with a Youtube Application API key, you won't be able to log into your account locally using OAuth. The tool will look for the environment variable `YOUTUBE_API_KEY` which should be set to the aforementioned key. The tool only accesses public data so you don't need any specfic permissions or anything like that.

Google's documentation outlines how to create the credentials but the gist of it is:

1) Create a new project in the Google Cloud Platform console.
2) Under `APIs & Services`, enable the YouTube Data API v3 for that project.
3) Under `APIs & Services > Credentials`, create credentials and select API key.

Build the binary.

`go build cmd/yougo-premium/main.go -o yougo`

Call the help command for usage information.

`./yougo help`

You can call `./yougo help <command>` to get command specific help.

## Demo

![Yougo-Premium Demo](gif/demo.gif)

## Current Features

* Streaming video download from followed channels
* Ability to add subscriptions via channel url or video url
* Tracks last refresh timestamp to only pull new videos
