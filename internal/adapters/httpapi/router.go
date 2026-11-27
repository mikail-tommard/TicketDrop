package httpapi

import (
	"net/http"
	"time"

	"github.com/mikail-tommard/TicketDrop/internal/domain"
	"github.com/mikail-tommard/TicketDrop/internal/usecase"
)

type Api struct {
	svc           *usecase.Service
	buyersByToken map[string]string
}

type ConcertResponce struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Available   int    `json:"available"`
}

type ReservationResponce struct {
	ID        string                   `json:"id"`
	Status    domain.ReservationStatus `json:"status"`
	ExpiresAt time.Time                `json:"expires_at"`
}

func New(svc *usecase.Service, buyersByToken map[string]string) *Api {
	return &Api{
		svc:           svc,
		buyersByToken: buyersByToken,
	}
}

func (a *Api) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /livez", a.HandlerLivez)
	mux.HandleFunc("GET /concerts/{id}", a.GetConcertByID)
	mux.HandleFunc("POST /concerts/{id}/reservations", a.CreateReservation)

	return mux
}
