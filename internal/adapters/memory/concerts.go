package memory

import (
	"context"

	"github.com/mikail-tommard/TicketDrop/internal/domain"
)

type ConcertRepo struct {
	concerts map[int]*domain.Concert
}

func NewMemory(concerts map[int]*domain.Concert) *ConcertRepo {
	return &ConcertRepo{concerts: concerts}
}

func (r *ConcertRepo) GetConcertByID(_ context.Context, id int) (*domain.Concert, error) {
	concert, ok := r.concerts[id]
	if !ok {
		return nil, domain.ErrConcertNotFound
	}
	return concert, nil
}
