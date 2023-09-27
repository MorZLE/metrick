package main

import (
	"github.com/MorZLE/metrick/config"
	"github.com/MorZLE/metrick/internal/server/handlers"
	"github.com/MorZLE/metrick/internal/server/services"
	"github.com/MorZLE/metrick/internal/server/storages"
)

func main() {
	cnf := config.NewConfigServer()

	repo := storages.NewStorage()
	logic := services.NewService(&repo)
	h := handlers.NewHandler(&logic, cnf)

	h.UpServer()
}
