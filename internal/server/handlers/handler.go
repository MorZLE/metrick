package handlers

import (
	"errors"
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

func (h *Handler) UpServer(port string) {
	h.routs(port)

}

func (h *Handler) routs(port string) {
	router := mux.NewRouter()
	router.HandleFunc(`/update/{metric}/{name}/{value}`, h.UpdateMetric)
	router.HandleFunc(`/value/{metric}/{name}`, h.ValueMetric)
	router.HandleFunc(`/`, h.ValueMetrics)
	http.Handle("/", router)
	log.Println(http.ListenAndServe(port, router))
}

func (h *Handler) UpdateMetric(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
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

func (h *Handler) ValueMetric(res http.ResponseWriter, req *http.Request) {
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

func (h *Handler) ValueMetrics(res http.ResponseWriter, _ *http.Request) {
	metrics := h.Logic.GetAllMetrics()

	_, err := res.Write([]byte(metrics))
	if err != nil {
		return
	}

	res.WriteHeader(http.StatusOK)

}
