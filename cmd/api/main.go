package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/mikail-tommard/TicketDrop/internal/adapters/httpapi"
)

func main() {
	api := httpapi.New()

	server := &http.Server{
		Addr:              ":8080",
		Handler:           api.Routes(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("Server is starting in port: %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
