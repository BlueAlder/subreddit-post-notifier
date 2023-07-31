package subreddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

const baseURI = "https://reddit.com"

type SubredditMonitor struct {
	subreddit  string
	latestPost Post
}

type Post struct {
	Url   string
	Title string
}

// Creates a new [SubredditMonitor] which will manage the process of detecting new
// subreddit posts
func New(subreddit string) (*SubredditMonitor, error) {
	s := &SubredditMonitor{
		subreddit:  subreddit,
		latestPost: Post{},
	}

	data, err := s.getLatestSubredditPost()
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve initial reddit data: %w", err)
	}
	log.Info().Msg(fmt.Sprintf("Setting initial post to be: %s", data.Title))
	s.latestPost.Title = data.Title
	s.latestPost.Url = data.URL
	return s, nil
}

// Returns the latest post that has been detected
func (s *SubredditMonitor) GetLatestPost() Post {
	return s.latestPost
}

func (s *SubredditMonitor) getLatestSubredditPost() (*Data, error) {
	uri := fmt.Sprintf("%s/r/%s/new.json?sort=new`", baseURI, s.subreddit)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Non 200 response: %s", string(data)))
	}

	var result APIResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	if len(result.Data.Children) > 0 {
		return &result.Data.Children[0].Data, nil
	} else {
		fmt.Printf(string(data))
		return nil, errors.New("no children")
	}

}

// Checks for new posts in the subreddit and returns true
// when there has been a new post detected
func (s *SubredditMonitor) CheckForNewPosts() bool {
	log.Debug().Msg("Checking for new posts...")
	data, err := s.getLatestSubredditPost()
	if err != nil {
		log.Warn().Err(err).Msg("Unable to retrieve reddit data.")
		return false
	}

	if data.URL != s.latestPost.Url {
		log.Info().Msg("Found new post!")
		log.Info().Msg(fmt.Sprintf("New post title: %s", data.Title))
		s.latestPost = Post{data.URL, data.Title}
		return true
	}
	log.Debug().Msg("No new post found")
	return false

}
