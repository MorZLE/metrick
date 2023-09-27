package main

import (
	"github.com/MorZLE/metrick/config"
	"github.com/MorZLE/metrick/internal/client/handlers"
	"github.com/MorZLE/metrick/internal/client/services"
	"github.com/MorZLE/metrick/internal/client/storages"
	"log"
)

func main() {
	cnf := config.NewConfigAgent()

	repo := storages.NewStorage()
	h := handlers.NewHandler()
	logic := services.NewService(&repo, &h, cnf)

	// Add logging statements
	log.Println("Starting UpClient...")
	logic.UpClient()
	log.Println("UpClient complete.")
}
