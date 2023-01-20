package repository

import "time"

type Phrase struct {
	ID              int       `json:"id,omitempty"`
	SenderChannelId string    `json:"sender_chan_id,omitempty"`
	SenderId        string    `json:"sender_id,omitempty"`
	SenderName      string    `json:"sender_name,omitempty"`
	PhraseText      string    `json:"phrase,omitempty"`
	Reply           string    `json:"reply_id,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Phrases map[int]*Phrase

type PhraseRepository interface {
	Create(phrase *Phrase) (*Phrase, error)
	Get(id int) (*Phrase, error)
}

type TopElement struct {
	Name  string
	Count int
}
