package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/BlueAlder/reddit-newpost-notify/pkg/notifier"
	"github.com/BlueAlder/reddit-newpost-notify/pkg/subreddit"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	subredditName := flag.String("subreddit", "", "The subreddit to monitor")
	pollingInterval := flag.Duration("interval", time.Second*10, "Interval to check for new posts in seconds")
	flag.Parse()

	if *subredditName == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Dev output for pretty
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	startPolling(*subredditName, *pollingInterval)

}

func startPolling(subredditName string, pollingInterval time.Duration) {
	log.Info().Msg(fmt.Sprintf("Starting monitoring subreddit /r/%s every %s", subredditName, pollingInterval))

	rm, err := subreddit.New(subredditName)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to start monitoring subreddit")
	}

	poll := time.NewTicker(pollingInterval)

	for range poll.C {
		fmt.Printf("\rLast Checked %s", time.Now().Format("15:04:05"))
		if rm.CheckForNewPosts() {
			notifier.OpenURLInBrowser(rm.GetLatestPost().Permalink)
		}
	}
}
