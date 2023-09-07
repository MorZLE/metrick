package main

import (
	"fmt"
	"github.com/MorZLE/metrick/internal/handlers"
	"github.com/MorZLE/metrick/internal/services"
	"github.com/MorZLE/metrick/internal/storages"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	repo := storages.Storage{}
	logic := services.Service{Storage: repo}
	h := handlers.Handler{Logic: logic}
	UpServer(h)
}

func UpServer(h handlers.Handler) {
	router := mux.NewRouter()
	router.HandleFunc(`/update/{metric}/{name}/{value}`, h.UpdatePage)
	http.Handle("/", router)

	err := http.ListenAndServe(`:8080`, nil)
	if err == nil {
		panic(fmt.Errorf("ошибка запуска сервера: %d", err))
	}

}
