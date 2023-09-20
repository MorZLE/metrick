package main

import (
	"github.com/MorZLE/metrick/internal/client/handlers"
	"github.com/MorZLE/metrick/internal/client/services"
	"github.com/MorZLE/metrick/internal/client/storages"
)

func main() {
	repo := storages.NewStorage()
	h := handlers.NewHandler()
	logic := services.NewService(&repo, &h)

	logic.UpClient()
}
