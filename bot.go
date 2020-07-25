package main

import (
	"log"

	"github.com/pkg/errors"
	"github.com/turnage/graw/reddit"
)

// Deogracias is the main handler for reddit events
type Deogracias struct {
	bot      reddit.Bot
	username string
}

func newDeogracias(bot reddit.Bot, username string) *Deogracias {
	return &Deogracias{
		bot:      bot,
		username: username,
	}
}

// Post receives a new event when a post is made to a subreddit that Deogracias is assigned to
func (d *Deogracias) Post(p *reddit.Post) error {
	if err := d.bot.Reply(p.Name, d.getPostQuote()); err != nil {
		log.Println(errors.WithStack(errors.Errorf("failed to make post reply: %v", err)))
	}
	return nil
}

// CommentReply receives a new event when a user replies to Deogracias
func (d *Deogracias) CommentReply(m *reddit.Message) error {
	if err := d.bot.Reply(m.Name, d.getReplyQuote()); err != nil {
		log.Println(errors.WithStack(errors.Errorf("failed to reply back: %v", err)))
	}
	return nil
}
