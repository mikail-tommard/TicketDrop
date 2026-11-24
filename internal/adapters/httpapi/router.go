package httpapi

import "net/http"

type Api struct{}

func New() *Api {
	return &Api{}
}

func (a *Api) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /livez", a.HandlerLivez)

	return mux
}
