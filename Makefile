BINARY=subreddit-post-notifier

build:
	go build -o .bin/ ./cmd/subreddit-post-notifier

build-windows:
	GOARCH=amd64 GOOS=windows go build -o .bin/ ./cmd/subreddit-post-notifier