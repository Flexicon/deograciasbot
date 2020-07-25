package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

func main() {
	bot, username, err := newRedditBot()
	if err != nil {
		log.Fatalln("Failed to create bot handle: ", err)
	}

	handler := newDeogracias(bot, username)
	cfg := graw.Config{
		Subreddits:     []string{"flexicondev", os.Getenv("SUBREDDITS")},
		CommentReplies: true,
	}

	_, wait, err := graw.Run(handler, bot, cfg)
	if err != nil {
		log.Fatalln("Failed to start graw run: ", err)
	}

	log.Println("Awaiting new acts of the Grievous Miracle...")
	log.Fatalln(wait())
}

func newRedditBot() (reddit.Bot, string, error) {
	username := os.Getenv("CLIENT_USERNAME")

	bot, err := reddit.NewBot(reddit.BotConfig{
		Agent: os.Getenv("USER_AGENT"),
		App: reddit.App{
			ID:       os.Getenv("CLIENT_ID"),
			Secret:   os.Getenv("CLIENT_SECRET"),
			Username: username,
			Password: os.Getenv("CLIENT_PASSWORD"),
		},
		Rate: 0,
	})

	return bot, username, err
}