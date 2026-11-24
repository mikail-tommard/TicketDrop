package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/mikail-tommard/TicketDrop/internal/adapters/httpapi"
	"github.com/mikail-tommard/TicketDrop/internal/adapters/memory"
	"github.com/mikail-tommard/TicketDrop/internal/domain"
	"github.com/mikail-tommard/TicketDrop/internal/usecase"
)

func main() {
	conserts := make(map[int]*domain.Concert, 0)

	concert, err := domain.NewConcert(1, "Вечер джаза", "Концерт в пятницу")
	if err != nil {
		log.Fatalf("не удалось создать демоконцерт: %v", err)
	}

	conserts[concert.ID()] = concert

	repo := memory.NewMemory(conserts)

	svc := usecase.NewService(repo)

	api := httpapi.New(svc)

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
