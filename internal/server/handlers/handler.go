package handlers

import (
	"errors"
	"github.com/MorZLE/metrick/config"
	"github.com/MorZLE/metrick/internal/logger"
	"github.com/MorZLE/metrick/internal/server/services"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func NewHandler(l services.ServiceInterface, cnf *config.ConfigServer) Handler {
	return Handler{Logic: l, cnf: cnf}
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=HandlerServer
type HandlerServer interface {
	UpServer()
	routs()
	UpdateMetric(res http.ResponseWriter, req *http.Request)
}

type Handler struct {
	HandlerServer
	Logic services.ServiceInterface
	cnf   *config.ConfigServer
}

func (h *Handler) UpServer() {
	logger.Initialize()

	router := mux.NewRouter()
	router.Handle(`/update/{metric}/{name}/{value}`, logger.RequestLogger(h.UpdateMetric))
	router.Handle(`/value/{metric}/{name}`, logger.RequestLogger(h.ValueMetric))
	router.Handle(`/`, logger.RequestLogger(h.ValueMetrics))

	logger.Log.Info("Running server", zap.String("address", h.cnf.FlagRunAddr))
	log.Println(http.ListenAndServe(h.cnf.FlagRunAddr, router))

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
