package main
/**

  1. build my own timeline information
 https://dev.twitter.com/overview/general
 2. Use import-io to competitors timeline
 https://api.import.io/store/connector/584db565-ca79-4ea9-99a7-04fabbe4a1d8/_query?input=webpage/url:https%3A%2F%2Ftwitter.com%2Floggly&&_apikey=b789e0e68a4742ae9371404b7f42f29c0804596902922ca914179ee50801d07d665c8ac953ecb29978051618bc5a939a66c2a4fe44b25a431cdbec258bbe90af041b8b87d229ccb1226b9c8add25cdfa

 */

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func printTweets(msg string, tweets []twitter.Tweet) {
	for _, value := range tweets {
		fmt.Println("Created:", value.CreatedAt, "Favourite:", value.FavoriteCount, "Retweet:", value.RetweetCount, " Msg:", value.Text)
	}
}

func main() {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "xaRcbnR7fuTrLYgL1u5nuHh3QN", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "IEcRkZnQbGEFNH5QLbNIM9dSEo5jYwiEgmLQcu1Jdy5vbyt5djE", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "315084325-bSOcspNUdmWWxENXhmtw27z8oMfFlH63cGVYwBylI", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "vr05niG62rJ2cNuTRyhZLeOpENgE4QXLQ9gJL14T1kOWuL", "Twitter Access Secret")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{SkipStatus: twitter.Bool(true)}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("\n\n\nUser's ACCOUNT:\n%+v\n", user)

	// Home Timeline
	homeTimelineParams := &twitter.HomeTimelineParams{Count: 2}
	tweets, _, _ := client.Timelines.HomeTimeline(homeTimelineParams)
	fmt.Printf("\n\n\nUser's HOME TIMELINE:\n%+v\n", tweets)

	// Mention Timeline
	mentionTimelineParams := &twitter.MentionTimelineParams{Count: 2}
	tweets, _, _ = client.Timelines.MentionTimeline(mentionTimelineParams)
	fmt.Printf("\n\n\nUser's MENTION TIMELINE:\n%+v\n", tweets)

	// Retweets of Me Timeline
	retweetTimelineParams := &twitter.RetweetsOfMeTimelineParams{Count: 10}
	tweets, _, _ = client.Timelines.RetweetsOfMeTimeline(retweetTimelineParams)

	printTweets("Retweets", tweets);
	//fmt.Printf("\n\n\nUser's 'RETWEETS OF ME' TIMELINE:\n%+v\n", tweets)

	// Update (POST!) Tweet (uncomment to run)
	// tweet, _, _ := client.Statuses.Update("just setting up my twttr", nil)
	// fmt.Printf("Posted Tweet\n%v\n", tweet)
}