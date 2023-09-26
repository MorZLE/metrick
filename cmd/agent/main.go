package main

import (
	"github.com/MorZLE/metrick/cmd/config"
	"github.com/MorZLE/metrick/internal/client/handlers"
	"github.com/MorZLE/metrick/internal/client/services"
	"github.com/MorZLE/metrick/internal/client/storages"
	"log"
)

func main() {
	config.ParseFlagsAgent()
	repo := storages.NewStorage()
	h := handlers.NewHandler()
	logic := services.NewService(&repo, &h, config.FlagPollInterval, config.FlagReportInterval, config.FlagAddr)

	// Add logging statements
	log.Println("Starting UpClient...")
	logic.UpClient()
	log.Println("UpClient complete.")
}
