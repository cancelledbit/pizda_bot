package repository

import (
	"context"
	"database/sql"
	"time"
)

type MysqlHagzorPhrasesRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewMysqlHagzorPhrasesRepository(ctx context.Context, db *sql.DB) *MysqlHagzorPhrasesRepository {
	return &MysqlHagzorPhrasesRepository{
		db: db, ctx: ctx,
	}
}

func (r MysqlHagzorPhrasesRepository) Get() (*HagzorPhrase, error) {
	query := "SELECT id, text FROM hagzor_phrases ORDER BY RAND() LIMIT 1;"

	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	var phrase HagzorPhrase
	err := r.fetchRow(r.db.QueryRowContext(ctx, query), &phrase)
	if err != nil {
		return nil, err
	}

	return &phrase, nil
}

func (r MysqlHagzorPhrasesRepository) Create(phrase *HagzorPhrase) (*HagzorPhrase, error) {
	ctx, cancel := context.WithTimeout(r.ctx, 15*time.Second)
	defer cancel()

	query := `INSERT INTO hagzor_phrases (text)
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

func (r *MysqlHagzorPhrasesRepository) fetchRow(row *sql.Row, phrase *HagzorPhrase) error {
	return row.Scan(
		&phrase.ID,
		&phrase.Text,
	)
}
