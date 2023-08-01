# Subreddit Post Notifier

Simple project that polls the reddit json API to check for new posts in a subreddit and will notify the user with their configured option.

## Usage

```bash
./subreddit-post-notifier --subreddit funny --interval 10s
```

## Notification options

Currently very limited, as in it will open the URL in the browser and send the
os a notification. I haven't made any mechanisms to control these either. But
here are some fun ideas I have for the future.

- Open the URL in a browser ✅
- Send a notification to the OS✅
- Send an email
- Send a webhook
- Send a text?
- Send a letter in the post