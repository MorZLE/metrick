package server

import (
	"fmt"
	"github.com/MorZLE/metrick/internal/services/server"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandler(l server.Service) Handler {
	return Handler{Logic: l}
}

type Handler struct {
	Logic server.Service
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

func (h Handler) New(s server.Service) *Handler {
	return &Handler{Logic: s}
}

func (h Handler) UpdatePage(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	metric := vars["metric"]
	name := vars["name"]
	value := vars["value"]
	if metric == "" {
		res.WriteHeader(http.StatusNotFound)
	}
	err := h.Logic.VerifType(metric, name, value)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
	}
	res.WriteHeader(http.StatusOK)
}
