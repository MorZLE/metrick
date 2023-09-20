package handlers

import (
	"errors"
	"fmt"
	"github.com/MorZLE/metrick/internal/server/services"
	"github.com/go-chi/chi/v5"
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
	err := http.ListenAndServe(`:8080`, nil)
	if err == nil {
		panic(fmt.Errorf("ошибка запуска сервера: %d", err))
	}
}

func (h Handler) routs() {
	router := chi.NewRouter()
	router.Post(`/update/{metric}/{name}/{value}`, h.UpdateMetric)
	router.Get(`/update/{metric}/{name}`, h.ValueMetric)
	router.Get(`/`, h.ValueMetrics)
	http.Handle("/", router)
}

func (h Handler) UpdateMetric(res http.ResponseWriter, req *http.Request) {
	var vars map[string]string

	vars["metric"] = chi.URLParam(req, "metric")
	vars["name"] = chi.URLParam(req, "name")
	vars["value"] = chi.URLParam(req, "value")

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
	var vars map[string]string

	vars["metric"] = chi.URLParam(req, "metric")
	vars["name"] = chi.URLParam(req, "name")
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
