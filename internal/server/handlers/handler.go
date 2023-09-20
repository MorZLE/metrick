package handlers

import (
	"errors"
	"fmt"
	"github.com/MorZLE/metrick/internal/server/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewHandler(l services.ServiceInterface) Handler {
	return Handler{Logic: l}
}

type HandlerServer interface {
	UpServer()
	routs()
	UpdateMetric(res http.ResponseWriter, req *http.Request)
}

type Handler struct {
	HandlerServer
	Logic services.ServiceInterface
}

func (h Handler) UpServer() {
	h.routs()

}

func (h Handler) routs() {
	router := mux.NewRouter()
	router.HandleFunc(`/update/{metric}/{name}/{value}`, h.UpdateMetric)
	router.HandleFunc(`/update/{metric}/{name}`, h.ValueMetric)
	router.HandleFunc(`/`, h.ValueMetrics)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func (h Handler) UpdateMetric(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Println(vars)
	err := h.Logic.ProcessingMetric(vars)
	if err != nil {
		if errors.Is(err, services.ErrBadRequest) {
			http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		}
		if errors.Is(err, services.ErrStatusNotFound) {
			http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (h Handler) ValueMetric(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	value, err := h.Logic.ValueMetric(vars)

	if err != nil {
		if errors.Is(err, services.ErrStatusNotFound) {
			http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
		return
	}

	_, err = res.Write([]byte(value))
	if err != nil {
		return
	}

	res.WriteHeader(http.StatusOK)

}

func (h Handler) ValueMetrics(res http.ResponseWriter, _ *http.Request) {
	metrics := h.Logic.GetAllMetrics()

	_, err := res.Write([]byte(metrics))
	if err != nil {
		return
	}

	res.WriteHeader(http.StatusOK)

}
