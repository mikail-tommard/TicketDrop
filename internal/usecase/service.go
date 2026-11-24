package usecase

import (
	"context"

	"github.com/mikail-tommard/TicketDrop/internal/domain"
)

type Repo interface {
	GetConcertByID(ctx context.Context, id int) (*domain.Concert, error)
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
