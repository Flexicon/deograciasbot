package main

import (
	"log"

	"github.com/pkg/errors"
	"github.com/turnage/graw/reddit"
)

// Deogracias is the main handler for reddit events
type Deogracias struct {
	bot reddit.Bot
}

func newDeogracias(bot reddit.Bot) *Deogracias {
	return &Deogracias{
		bot: bot,
	}
}

// Post receives a new event when a post is made to a subreddit that Deogracias is assigned to
func (d *Deogracias) Post(p *reddit.Post) error {
	err := d.bot.Reply(p.Name, d.getPostQuote())
	if err != nil {
		log.Println(errors.WithStack(errors.Errorf("failed to make post reply: %v", err)))
	}
	return nil
}

// CommentReply receives a new event when a user replies to Deogracias
func (d *Deogracias) CommentReply(m *reddit.Message) error {
	err := d.bot.Reply(m.Name, d.getReplyQuote())
	if err != nil {
		log.Println(errors.WithStack(errors.Errorf("failed to reply back: %v", err)))
	}
	return nil
}
