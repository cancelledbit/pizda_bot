package repository

import (
	"context"
	"database/sql"
	"time"
)

type MysqlEUPhrasesRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewMysqlEUPhrasesRepository(ctx context.Context, db *sql.DB) *MysqlEUPhrasesRepository {
	return &MysqlEUPhrasesRepository{
		db: db, ctx: ctx,
	}
}

func (r MysqlEUPhrasesRepository) Get() (*EUPhrase, error) {
	query := "SELECT id, text FROM eu_phrases ORDER BY RAND() LIMIT 1;"

	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	var phrase EUPhrase
	err := r.fetchRow(r.db.QueryRowContext(ctx, query), &phrase)
	if err != nil {
		return nil, err
	}
	return &phrase, nil
}

func (r MysqlEUPhrasesRepository) Create(phrase *EUPhrase) (*EUPhrase, error) {
	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	query := `INSERT INTO eu_phrases (text)
		VALUES (?)
		RETURNING id, text;`
	row := r.db.QueryRowContext(
		ctx,
		query,
		phrase.Text,
	)
	err := r.fetchRow(row, phrase)
	if err != nil {
		return nil, err
	}
	return phrase, nil
}

func (r *MysqlEUPhrasesRepository) fetchRow(row *sql.Row, phrase *EUPhrase) error {
	return row.Scan(
		&phrase.ID,
		&phrase.Text,
	)
}
