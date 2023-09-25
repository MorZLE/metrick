package main

import (
	"github.com/MorZLE/metrick/cmd/flags"
	"github.com/MorZLE/metrick/internal/client/handlers"
	"github.com/MorZLE/metrick/internal/client/services"
	"github.com/MorZLE/metrick/internal/client/storages"
	"log"
)

func main() {
	flags.ParseFlagsAgent()
	repo := storages.NewStorage()
	h := handlers.NewHandler()
	logic := services.NewService(&repo, &h, flags.FlagPollInterval, flags.FlagReportInterval, flags.FlagAddr)

	// Add logging statements
	log.Println("Starting UpClient...")
	logic.UpClient()
	log.Println("UpClient complete.")
}
