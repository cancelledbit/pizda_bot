package repository

import "time"

type Channel struct {
	ID          int       `json:"id,omitempty"`
	ChannelId   string    `json:"channel_id,omitempty"`
	ChannelName string    `json:"channel_name,omitempty"`
	Type        string    `json:"type,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ChannelRepository interface {
	Create(channel *Channel) (*Channel, error)
	Get(id int) (*Channel, error)
}
