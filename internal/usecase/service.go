package usecase

import (
	"context"
	"strings"
	"time"

	"github.com/mikail-tommard/TicketDrop/internal/domain"
)

type Repo interface {
	GetConcertByID(ctx context.Context, id int) (*domain.Concert, error)

	CreateReservation(ctx context.Context, reservation *domain.Reservation, key string) (*domain.Reservation, error)
	// GetReservationByID(cxt context.Context, id string) (*domain.Reservation, error)
}

type ReservationInput struct {
	ID        string
	ConcertID int
	BuyerID   string
	Status    domain.ReservationStatus
	ExpiresAt time.Time
}

type Service struct {
	repo Repo
}

func NewService(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetConcertByID(ctx context.Context, id int) (*domain.Concert, error) {
	return s.repo.GetConcertByID(ctx, id)
}

func (s *Service) CreateReservation(ctx context.Context, reservation ReservationInput, key string) (*domain.Reservation, error) {
	key = strings.TrimSpace(key)
	if key == "" {
		return nil, domain.ErrKey
	}

	reservation.ExpiresAt = time.Now().Add(2 * time.Minute)

	booking, err := domain.NewReservation("", reservation.BuyerID, reservation.ConcertID, reservation.ExpiresAt)
	if err != nil {
		return nil, err
	}

	res, err := s.repo.CreateReservation(ctx, booking, key)
	if err != nil {
		return nil, err
	}

	return res, nil
}
