package handlers

import (
	"fmt"
	"github.com/MorZLE/metrick/internal/server/services"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandler(l services.Service) Handler {
	return Handler{Logic: l}
}

type Handler struct {
	Logic services.Service
}

func (h Handler) UpServer() {
	h.routs()
	err := http.ListenAndServe(`:8080`, nil)
	if err == nil {
		panic(fmt.Errorf("ошибка запуска сервера: %d", err))
	}
}

func (h Handler) routs() {
	router := mux.NewRouter()
	router.HandleFunc(`/update/{metric}/{name}/{value}`, h.UpdatePage)
	http.Handle("/", router)
}

func (h Handler) UpdatePage(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	err := h.Logic.ProcessingMetrick(vars)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	res.WriteHeader(http.StatusOK)
}
