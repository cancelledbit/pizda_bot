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

func (r *MysqlChannelRepository) Create(channel *Channel) (*Channel, error) {
	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	query := `INSERT INTO channels (channel_id, channel_name, type)
		VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE channel_name = ? 
		RETURNING id, channel_id, channel_name, type, created_at, updated_at;`
	row := r.db.QueryRowContext(ctx, query, channel.ChannelId, channel.ChannelName, channel.Type, channel.ChannelName)
	err := r.fetchRow(row, channel)
	if err != nil {
		return nil, err
	}
	return channel, nil
}

func (r *MysqlChannelRepository) Get(id int) (*Channel, error) {
	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()
	var channel Channel

	query := "SELECT id, channel_id, channel_name, type, created_at, updated_at FROM channels WHERE id = $1"
	err := r.fetchRow(r.db.QueryRowContext(ctx, query, id), &channel)
	if err != nil {
		return nil, err
	}
	return &channel, nil
}

func (r *MysqlChannelRepository) GetByOffset(offset int, limit int) (*Channels, error) {
	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	query := "SELECT * FROM channels LIMIT ?, ?"
	rows, err := r.db.QueryContext(ctx, query, offset, limit)
	if err != nil {
		return nil, err
	}
	channels := make(Channels)
	for rows.Next() {
		var channel Channel
		err := r.fetchRows(rows, &channel)
		if err != nil {
			return nil, err
		}
		channels[channel.ID] = &channel
	}
	return &channels, nil
}

func (r *MysqlChannelRepository) fetchRow(row *sql.Row, channel *Channel) error {
	return row.Scan(&channel.ID, &channel.ChannelId, &channel.ChannelName, &channel.Type, &channel.CreatedAt, &channel.UpdatedAt)
}

func (r *MysqlChannelRepository) fetchRows(rows *sql.Rows, channel *Channel) error {
	return rows.Scan(&channel.ID, &channel.ChannelId, &channel.ChannelName, &channel.Type, &channel.CreatedAt, &channel.UpdatedAt)
}
