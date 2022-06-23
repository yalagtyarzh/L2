package main

import (
	"log"
	"net/http"

	"dev11/config"
	"dev11/handlers"
	"dev11/logging"
	"dev11/middleware"
	"dev11/repository/dbrepo"
	"dev11/router"
)

func main() {
	log.Println("reading environment")
	cfg := config.GetConfig()

	log.Println("router initializing")
	mux := router.New()

	log.Println("logger initializing")
	logger := logging.New()

	log.Println("database initializing")
	repo := dbrepo.NewMemoryStorage()

	h := handlers.NewRepo(repo)
	handlers.NewHandler(h)

	handler := middleware.EventLogger(mux, logger)

	log.Println("Starting!")
	log.Fatal(http.ListenAndServe(cfg.IP+cfg.Port, handler))
}
