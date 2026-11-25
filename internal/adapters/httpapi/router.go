package httpapi

import (
	"net/http"

	"github.com/mikail-tommard/TicketDrop/internal/usecase"
)

type Api struct {
	svc *usecase.Service
}

type ConcertResponce struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Available   int    `json:"available"`
}

func New(svc *usecase.Service) *Api {
	return &Api{
		svc: svc,
	}
}

func (a *Api) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /livez", a.HandlerLivez)
	mux.HandleFunc("GET /concerts/{id}", a.GetConcertByID)

	return mux
}
