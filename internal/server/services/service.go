package services

import (
	"errors"
	"github.com/MorZLE/metrick/internal/server"
	"github.com/MorZLE/metrick/internal/server/storages"
	"strconv"
)

var ErrBadRequest error = errors.New("BadRequest")
var ErrStatusNotFound error = errors.New("StatusNotFound")

func NewService(s storages.MemStorage) Service {
	return Service{Storage: s}
}

type Service struct {
	Storage storages.MemStorage
}

func (s Service) ProcessingMetrick(vars map[string]string) error {
	metric := vars["metric"]
	name := vars["name"]
	value := vars["value"]
	if metric == "" {
		return ErrStatusNotFound
	}
	if metric != "gauge" && metric != "counter" {
		return errors.New("http.StatusBadRequest")
	}
	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return ErrBadRequest
	}
	if metric != "counter" {
		s.Storage.AddCounter(server.Gauge{Metric: metric, Name: name, Value: valueFloat})
	}
	if metric != "gauge" {
		s.Storage.AddGauge(server.Gauge{Metric: metric, Name: name, Value: valueFloat})
	}
	return nil

}
