package main

import (
	"github.com/MorZLE/metrick/cmd/config"
	"github.com/MorZLE/metrick/internal/server/handlers"
	"github.com/MorZLE/metrick/internal/server/services"
	"github.com/MorZLE/metrick/internal/server/storages"
)

func main() {
	config.ParseFlagsServer()

	repo := storages.NewStorage()
	logic := services.NewService(&repo)
	h := handlers.NewHandler(&logic)

	h.UpServer(config.FlagRunAddr)
}
