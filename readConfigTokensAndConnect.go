package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

//Config stores the data from json twitterConfig.json file
type Config struct {
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessTokenKey    string `json:"access_token_key"`
	AccessTokenSecret string `json:"access_token_secret"`
}

func readConfigTokensAndConnect() (client *twitter.Client) {
	var config Config
	file, e := ioutil.ReadFile("twitterConfig.json")
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &config)
	fmt.Println("twitterConfig.json read comlete")

	fmt.Print("connecting to twitter api --> ")
	configu := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecret)
	token := oauth1.NewToken(config.AccessTokenKey, config.AccessTokenSecret)
	httpClient := configu.Client(oauth1.NoContext, token)
	// twitter client
	client = twitter.NewClient(httpClient)

	fmt.Println("connection successfull")

	return client
}
