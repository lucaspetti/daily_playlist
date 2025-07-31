# Daily Playlist

Get a random playlist from any user and open directly on your browser

## Getting started

This project relies on the [Go Spotify API Client](https://github.com/zmb3/spotify)

Copy the .env.example and set all needed env variables:

```bash
cp .env.example .env
```

ID and SECRET will be used to authenticate against the Spotify API.

The SPOTIFY_USER_ID will be used to get the given user playlists and open one of them in the browser.

Then just build the project with ```go build .``` and run with ```./daily_playlist```

## Notes

This script currently runs on Linux and opens the playlist url with google-chrome. Make sure to adapt the command for your operating system and preferred browser.
