package main

import (
	"context"
	"fmt"
	"os"
	"log"
	"os/exec"
	"math/rand"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	fmt.Println("Loading env...")

	userId := os.Getenv("SPOTIFY_USER_ID")
	if userId == "" {
		log.Fatalf("Missing user ID")
	}

	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("Could not get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)
	playlistPage, err := client.GetPlaylistsForUser(userId)

	if err != nil {
		log.Fatalf("Could not get user playlists: %v", err)
	}

	playlistsLength := len(playlistPage.Playlists)
	randomIndex := getRandomIndex(playlistsLength)

	pl := playlistPage.Playlists[randomIndex]
	fmt.Println("Chosen playlist: ", pl.Name)

	url := pl.ExternalURLs["spotify"]
	fmt.Println(url)

	openDailyPlaylist(url)
}

func getRandomIndex(length int) (n int) {
	rand.Seed(time.Now().Unix())
	n = rand.Int() % length
	return
}

func openDailyPlaylist(url string) {
	fmt.Println("Launching in chrome...")
	err := exec.Command("google-chrome", url).Start()

	if err != nil {
		fmt.Printf("Could not open daily playlist")
	}
}
