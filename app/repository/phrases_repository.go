package repository

import (
	"context"
	"database/sql"
	"time"
)

type MysqlPhrasesRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewMysqlPhrasesRepository(ctx context.Context, db *sql.DB) *MysqlPhrasesRepository {
	return &MysqlPhrasesRepository{
		db: db, ctx: ctx,
	}
}

func (r MysqlPhrasesRepository) Create(phrase *Phrase) (*Phrase, error) {
	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	query := `INSERT INTO phrases (sender_chan_id, sender_id, sender_name, phrase, reply_id)
		VALUES (?, ?, ?, ?, ?)
		RETURNING id, sender_chan_id, sender_id, sender_name, phrase, reply_id, created_at, updated_at;`
	err := r.db.QueryRowContext(
		ctx,
		query,
		phrase.SenderChannelId,
		phrase.SenderId,
		phrase.SenderName,
		phrase.PhraseText,
		phrase.Reply,
	).Scan(
		&phrase.ID,
		&phrase.SenderChannelId,
		&phrase.SenderId,
		&phrase.SenderName,
		&phrase.PhraseText,
		&phrase.Reply,
		&phrase.CreatedAt,
		&phrase.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return phrase, nil
}

func (r MysqlPhrasesRepository) Get(id int) (*Phrase, error) {
	query := "SELECT id, sender_chan_id, sender_id, sender_name, phrase, reply_id, created_at, updated_at FROM phrases WHERE id = $1"

	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()
	var phrase Phrase
	err := r.db.QueryRowContext(ctx, query, id).
		Scan(
			&phrase.ID,
			&phrase.SenderChannelId,
			&phrase.SenderId,
			&phrase.SenderName,
			&phrase.PhraseText,
			&phrase.Reply,
			&phrase.CreatedAt,
			&phrase.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &phrase, nil
}
