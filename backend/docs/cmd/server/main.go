package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"kindergarten-sun-backend/internal/handlers"
	"kindergarten-sun-backend/internal/service"
	"kindergarten-sun-backend/internal/storage"
)

func main() {
	store := storage.NewStore()
	svc := service.New(store)
	h := handlers.New(svc)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handlers.WithCORS(mux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Backend запущен: http://localhost:" + port)
	log.Println("Swagger UI: http://localhost:" + port + "/swagger/")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
