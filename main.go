package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var (
	consumerKey    = os.Getenv("CONSUMER_KEY")
	consumerSecret = os.Getenv("CONSUMER_SECRET")
	accessToken    = os.Getenv("ACCESS_TOKEN")
	accessSecret   = os.Getenv("ACCESS_SECRET")
)

func main() {
	api := anaconda.NewTwitterApiWithCredentials(accessToken, accessSecret, consumerKey, consumerSecret)
	homeTimelineResult, _ := api.GetHomeTimeline(nil)
	fmt.Println(homeTimelineResult)
}
