package domain

import "time"

type ReservationStatus string

const (
	ReservationHeld      ReservationStatus = "held"
	ReservationConfirmed ReservationStatus = "confirmed"
	ReservationExpired   ReservationStatus = "expired"
)

type Reservation struct {
	id        string
	concertID int
	buyerID   string
	status    ReservationStatus
	expiresAt time.Time
}

func NewReservation(id, buyerID string, concertID int, expiresAt time.Time) (*Reservation, error) {
	return &Reservation{
		id:        id,
		buyerID:   buyerID,
		concertID: concertID,
		status:    ReservationHeld,
		expiresAt: expiresAt,
	}, nil
}

func (r *Reservation) ID() string {
	return r.id
}
func (r *Reservation) BuyerID() string {
	return r.buyerID
}
func (r *Reservation) ConcertID() int {
	return r.concertID
}
func (r *Reservation) Status() ReservationStatus {
	return r.status
}
func (r *Reservation) ExpiresAt() time.Time {
	return r.expiresAt
}
