package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mikail-tommard/TicketDrop/internal/domain"
)

type Repo struct {
	db *sql.DB
}

func NewDB(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetConcertByID(ctx context.Context, id int) (*domain.Concert, error) {
	const q = `
		SELECT id, title, description, capacity, available
		FROM concerts
		WHERE id = $1
	`

	var (
		concertID   int
		title       string
		description string
		available   int
		capacity    int
	)

	err := r.db.QueryRowContext(ctx, q, id).Scan(&concertID, &title, &description, &capacity, &available)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrConcertNotFound
		}
		return nil, err
	}
	return domain.FromStorage(concertID, capacity, available, title, description), nil
}
