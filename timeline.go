package main

import (
	"github.com/dghubble/go-twitter/twitter"
)

func getTimeline(client *twitter.Client) []twitter.Tweet {
	homeTimelineParams := &twitter.HomeTimelineParams{
		Count: 4,
	}
	tweets, _, _ := client.Timelines.HomeTimeline(homeTimelineParams)
	return (tweets)
}
