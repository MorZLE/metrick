package handlers

import (
	"fmt"
	"github.com/MorZLE/metrick/internal/services"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	Logic services.Service
}

func (h Handler) New(s services.Service) *Handler {
	return &Handler{Logic: s}
}

func (h Handler) UpdatePage(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	metric := vars["metric"]

	if metric != "gauge" || metric != "counter" {
		res.WriteHeader(http.StatusBadRequest)
	}
	name := vars["name"]
	value := vars["value"]
	fmt.Println(metric, name, value)
}
