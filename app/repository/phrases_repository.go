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
	row := r.db.QueryRowContext(
		ctx,
		query,
		phrase.SenderChannelId,
		phrase.SenderId,
		phrase.SenderName,
		phrase.PhraseText,
		phrase.Reply,
	)
	err := r.fetchRow(row, phrase)
	if err != nil {
		return nil, err
	}
	return phrase, nil
}

func (r MysqlPhrasesRepository) Get(id int) (*Phrase, error) {
	query := "SELECT id, sender_chan_id, sender_id, sender_name, phrase, reply_id, created_at, updated_at FROM phrases WHERE id = ?"

	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	var phrase Phrase
	err := r.fetchRow(r.db.QueryRowContext(ctx, query, id), &phrase)
	if err != nil {
		return nil, err
	}
	return &phrase, nil
}

func (r MysqlPhrasesRepository) GetPhrasesByUserId(senderId string) (*Phrases, error) {
	query := "SELECT * FROM phrases WHERE sender_id = ?"

	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, senderId)
	if err != nil {
		return nil, err
	}
	phrases := make(Phrases)
	for rows.Next() {
		var phrase Phrase
		err := r.fetchRows(rows, &phrase)
		if err != nil {
			return nil, err
		}
		phrases[phrase.ID] = &phrase
	}
	return &phrases, nil
}

func (r *MysqlPhrasesRepository) GetByOffset(offset int, limit int) (*Phrases, error) {
	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	query := "SELECT * FROM phrases LIMIT ?, ?"
	rows, err := r.db.QueryContext(ctx, query, offset, limit)
	if err != nil {
		return nil, err
	}
	phrases := make(Phrases)
	for rows.Next() {
		var phrase Phrase
		err := r.fetchRows(rows, &phrase)
		if err != nil {
			return nil, err
		}
		phrases[phrase.ID] = &phrase
	}
	return &phrases, nil
}

func (r *MysqlPhrasesRepository) fetchRow(row *sql.Row, phrase *Phrase) error {
	return row.Scan(
		&phrase.ID,
		&phrase.SenderChannelId,
		&phrase.SenderId,
		&phrase.SenderName,
		&phrase.PhraseText,
		&phrase.Reply,
		&phrase.CreatedAt,
		&phrase.UpdatedAt,
	)
}

func (r *MysqlPhrasesRepository) fetchRows(rows *sql.Rows, phrase *Phrase) error {
	return rows.Scan(
		&phrase.ID,
		&phrase.SenderChannelId,
		&phrase.SenderId,
		&phrase.SenderName,
		&phrase.PhraseText,
		&phrase.Reply,
		&phrase.CreatedAt,
		&phrase.UpdatedAt,
	)
}
