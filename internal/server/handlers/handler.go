package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/MorZLE/metrick/config"
	"github.com/MorZLE/metrick/internal/constants"
	"github.com/MorZLE/metrick/internal/logger"
	"github.com/MorZLE/metrick/internal/server/services"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strconv"
)

func NewHandler(l services.ServiceInterface, cnf *config.ConfigServer) Handler {
	return Handler{logic: l, cnf: cnf}
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=HandlerServer
type HandlerServer interface {
	UpServer()
	routs()

	UpdateMetricJSON(res http.ResponseWriter, req *http.Request)
	ValueMetricJSON(res http.ResponseWriter, req *http.Request)

	ValueMetrics(res http.ResponseWriter, req *http.Request)
	UpdateMetric(res http.ResponseWriter, req *http.Request)
	ValueMetric(res http.ResponseWriter, req *http.Request)

	ResponseValueJSON(res http.ResponseWriter, metric, name string)
}

type Handler struct {
	logic services.ServiceInterface
	cnf   *config.ConfigServer
}

func (h *Handler) UpServer() {
	logger.Initialize()

	router := mux.NewRouter()
	router.Handle(`/update`, logger.RequestLogger(h.UpdateMetricJSON))
	router.Handle(`/value`, logger.RequestLogger(h.ValueMetricJSON))
	router.Handle(`/update/{metric}/{name}/{value}`, logger.RequestLogger(h.UpdateMetric))
	router.Handle(`/value/{metric}/{name}`, logger.RequestLogger(h.ValueMetric))
	router.Handle(`/`, logger.RequestLogger(h.ValueMetrics))

	logger.Log.Info("Running server", zap.String("address", h.cnf.FlagRunAddr))
	log.Println(http.ListenAndServe(h.cnf.FlagRunAddr, router))

}

func (h *Handler) UpdateMetric(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	metric := vars["metric"]
	name := vars["name"]
	value := vars["value"]

	err := h.logic.ProcessingMetric(metric, name, value)
	if err != nil {
		if errors.Is(err, constants.ErrBadRequest) {
			http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		}
		if errors.Is(err, constants.ErrStatusNotFound) {
			http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	res.WriteHeader(http.StatusOK)
}

func (h *Handler) ValueMetric(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	metric := vars["metric"]
	name := vars["name"]
	value, err := h.logic.ValueMetric(metric, name)

	if err != nil {
		if errors.Is(err, constants.ErrStatusNotFound) {
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

func (h *Handler) UpdateMetricJSON(res http.ResponseWriter, req *http.Request) {
	var metricJSON constants.Metrics
	var buf bytes.Buffer
	var value string
	// читаем тело запроса
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &metricJSON); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	metric := metricJSON.ID
	name := metricJSON.MType

	switch metricJSON.ID {
	case "gauge":
		value = strconv.FormatFloat(*metricJSON.Value, 'f', -1, 64)
	case "counter":
		value = strconv.FormatInt(*metricJSON.Delta, 10)
	default:
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	err = h.logic.ProcessingMetric(metric, name, value)

	if err != nil {
		if errors.Is(err, constants.ErrBadRequest) {
			http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		}
		if errors.Is(err, constants.ErrStatusNotFound) {
			http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	h.ResponseValueJSON(res, metric, name)

}

func (h *Handler) ValueMetricJSON(res http.ResponseWriter, req *http.Request) {
	var metricJSON constants.Metrics
	var buf bytes.Buffer

	// читаем тело запроса
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &metricJSON); err != nil {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	metric := metricJSON.ID
	name := metricJSON.MType
	h.ResponseValueJSON(res, metric, name)

}

func (h *Handler) ResponseValueJSON(res http.ResponseWriter, metric, name string) {

	obj, err := h.logic.ValueMetricJSON(metric, name)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	resp, err := json.Marshal(obj)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(resp)

}

func (h *Handler) ValueMetrics(res http.ResponseWriter, _ *http.Request) {
	metrics := h.logic.GetAllMetrics()

	_, err := res.Write([]byte(metrics))
	if err != nil {
		return
	}

	res.WriteHeader(http.StatusOK)

}
