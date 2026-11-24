package httpapi

import "net/http"

func (a *Api) HandlerLivez(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, 200, "Server is live")
}
