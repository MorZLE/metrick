package services

import (
	"errors"
	"github.com/MorZLE/metrick/internal/server"
	"github.com/MorZLE/metrick/internal/server/storages"
	"strconv"
)

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
		return errors.New("http.StatusNotFound")
	}
	if metric != "gauge" && metric != "counter" {
		return errors.New("http.StatusBadRequest")
	}
	valueFloat, err := strconv.ParseFloat(value, 16)
	if err != nil {
		return errors.New("http.StatusBadRequest")
	}
	if metric != "counter" {
		s.Storage.AddCounter(server.Gauge{Metric: metric, Name: name, Value: valueFloat})
	}
	if metric != "gauge" {
		s.Storage.AddGauge(server.Gauge{Metric: metric, Name: name, Value: valueFloat})
	}
	return nil

}
