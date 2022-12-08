package repository

import (
	"context"
	"database/sql"
	"time"
)

type MysqlChannelRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewMysqlChannelRepository(ctx context.Context, db *sql.DB) *MysqlChannelRepository {
	return &MysqlChannelRepository{
		db: db, ctx: ctx,
	}
}

func (r MysqlChannelRepository) Create(channel *Channel) (*Channel, error) {
	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	query := `INSERT INTO channels (channel_id, channel_name, type)
		VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE channel_name = ? 
		RETURNING id, channel_id, channel_name, type, created_at, updated_at;`
	err := r.db.QueryRowContext(ctx, query, channel.ChannelId, channel.ChannelName, channel.Type, channel.ChannelName).
		Scan(&channel.ID, &channel.ChannelId, &channel.ChannelName, &channel.Type, &channel.CreatedAt, &channel.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return channel, nil
}

func (r MysqlChannelRepository) Get(id int) (*Channel, error) {
	query := "SELECT id, channel_id, channel_name, type, created_at, updated_at FROM channels WHERE id = $1"

	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()
	var channel Channel
	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&channel.ID, &channel.ChannelId, &channel.ChannelName, &channel.Type, &channel.CreatedAt, &channel.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &channel, nil
}
