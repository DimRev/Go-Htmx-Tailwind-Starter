package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/DimRev/Go-Htmx-Tailwind-Starter/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Handle("/*", public())
	router.Get("/foo", handlers.Make(handlers.HandleHealthz))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server starter", "listen addr", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
