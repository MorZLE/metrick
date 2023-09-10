package main

import (
	server3 "github.com/MorZLE/metrick/internal/handlers/server"
	server2 "github.com/MorZLE/metrick/internal/services/server"
	"github.com/MorZLE/metrick/internal/storages/server"
)

func main() {

	repo := server.NewStorage()
	logic := server2.NewService(repo)
	h := server3.NewHandler(logic)
	h.UpServer()

}
