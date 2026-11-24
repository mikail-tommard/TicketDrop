package httpapi

import (
	"net/http"
	"strconv"
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
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := ConcertResponce{
		Title:       concert.Title(),
		Description: concert.Desc(),
	}

	writeJSON(w, http.StatusOK, resp)
}
