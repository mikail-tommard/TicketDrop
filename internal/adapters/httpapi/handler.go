package httpapi

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/mikail-tommard/TicketDrop/internal/domain"
	"github.com/mikail-tommard/TicketDrop/internal/usecase"
)

func (a *Api) HandlerLivez(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, 200, "Server is live")
}

func (a *Api) GetConcertByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		writeJSON(w, http.StatusBadRequest, "invalid request")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, "invalid request")
		return
	}

	concert, err := a.svc.GetConcertByID(r.Context(), i)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	resp := ConcertResponce{
		Title:       concert.Title(),
		Description: concert.Desc(),
		Available:   concert.Available(),
	}

	writeJSON(w, http.StatusOK, resp)
}

func (a *Api) CreateReservation(w http.ResponseWriter, r *http.Request) {
	parts := strings.Fields(r.Header.Get("Authorization"))
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		writeJSON(w, http.StatusUnauthorized, "invalud authorization")
		return
	}

	buyerID, ok := a.buyersByToken[parts[1]]
	if !ok {
		writeJSON(w, http.StatusUnauthorized, "invalid authorization")
		return
	}

	key := strings.TrimSpace(r.Header.Get("Idempotency-Key"))
	if key == "" {
		writeJSON(w, http.StatusBadRequest, "Idempotency-Key is required")
		return
	}

	concertID := r.PathValue("id")
	cID, err := strconv.Atoi(concertID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, "invalid request")
		return
	}

	res, err := a.svc.CreateReservation(r.Context(), usecase.ReservationInput{
		ConcertID: cID,
		BuyerID:   buyerID,
	}, key)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrSoldOut):
			writeJSON(w, http.StatusConflict, "Мест больше нет")
		case errors.Is(err, domain.ErrAlreadyBookerd):
			writeJSON(w, http.StatusConflict, "Ключ уже использован для другого концерта")
		default:
			writeJSON(w, http.StatusInternalServerError, "Не удалось создать бронь")
		}
		return
	}

	responce := ReservationResponce{
		ID:        res.ID(),
		Status:    res.Status(),
		ExpiresAt: res.ExpiresAt(),
	}

	writeJSON(w, http.StatusCreated, responce)

}
