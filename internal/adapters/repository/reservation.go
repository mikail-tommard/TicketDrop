package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/mikail-tommard/TicketDrop/internal/domain"
)

func (r *Repo) CreateReservation(ctx context.Context, reservation *domain.Reservation, key string) (*domain.Reservation, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var (
		id        string
		buyerID   string
		concertID int
		status    domain.ReservationStatus
		expiresAt time.Time
	)

	err = tx.QueryRowContext(ctx, `
		INSERT INTO reservations (concert_id, buyer_id, key, status, expires_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (buyer_id, key) DO NOTHING
		RETURNING id::text, buyer_id, concert_id, status, expires_at
	`, reservation.ConcertID(), reservation.BuyerID(), key, reservation.Status(), reservation.ExpiresAt()).Scan(&id, &buyerID, &concertID, &status, &expiresAt)

	if errors.Is(err, sql.ErrNoRows) {
		err = tx.QueryRowContext(ctx, `
        SELECT id::text, buyer_id, concert_id, status, expires_at
        FROM reservations
        WHERE buyer_id = $1 AND key = $2
    `, reservation.BuyerID(), key).
			Scan(&id, &buyerID, &concertID, &status, &expiresAt)
		if err != nil {
			return nil, err
		}

		if concertID != reservation.ConcertID() {
			return nil, domain.ErrAlreadyBookerd
		}

		return domain.FromStorageReservation(
			id, buyerID, status, concertID, expiresAt,
		), nil
	}
	if err != nil {
		return nil, err
	}

	results, err := tx.ExecContext(ctx, `
		UPDATE concerts
		SET available = available - 1
		WHERE id = $1 AND available > 0
	`, reservation.ConcertID())
	if err != nil {
		return nil, err
	}

	affected, err := results.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("no available seats")
	}

	return domain.FromStorageReservation(id, buyerID, status, concertID, expiresAt), tx.Commit()
}
