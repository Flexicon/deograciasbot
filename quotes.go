package main

import (
	"math/rand"
	"time"
)

// Slices containing all Deogracias quotes
var (
	GenericQuotes = []string{
		"For twisted are, were and will be, the paths of the Miracle.",
		"For fiery are the paths of the Miracle.",
		"And thus, guilt, repentance, mourning and every pain of the soul of all kind were visibly and tangibly manifested, everywhere and in all of us.",
		"Even a wise penitent like me knows nothing of what this is about.",
	}

	DirectedQuotes = []string{
		"Regretful be the heart, Penitent One.",
		"Sorrowful be the Heart, Penitent One.",
	}

	ReplyQuotes = append(GenericQuotes, DirectedQuotes...)
)

func (d *Deogracias) getRandomQuote(qoutes []string) string {
	rand.Seed(time.Now().Unix())
	return qoutes[rand.Intn(len(qoutes))]
}

func (d *Deogracias) getReplyQuote() string {
	return d.getRandomQuote(ReplyQuotes)
}

func (d *Deogracias) getPostQuote() string {
	return d.getRandomQuote(GenericQuotes)
}
