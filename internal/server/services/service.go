package services

import (
	"errors"
	"fmt"
	"github.com/MorZLE/metrick/internal/server"
	"github.com/MorZLE/metrick/internal/server/storages"
	"strconv"
	"strings"
)

var ErrBadRequest = errors.New("BadRequest")
var ErrStatusNotFound = errors.New("StatusNotFound")

func NewService(s storages.Repositories) Service {
	return Service{Storage: s}
}

type ServiceInterface interface {
	ProcessingMetric(vars map[string]string) error
	ValueMetric(vars map[string]string) (string, error)
	GetAllMetrics() string
}

type Service struct {
	ServiceInterface
	Storage storages.Repositories
}

func (s *Service) ProcessingMetric(vars map[string]string) error {
	metric := vars["metric"]
	name := vars["name"]
	value := vars["value"]
	if metric == "" {
		return ErrStatusNotFound
	}
	if metric != "gauge" && metric != "counter" {
		return ErrBadRequest
	}

	if metric == "counter" {
		value, err := strconv.Atoi(value)
		if err != nil {
			return ErrBadRequest
		}
		s.Storage.AddCounter(server.Counter{Metric: metric, Name: name, Value: value})
	}
	if metric == "gauge" {
		valueFloat, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return ErrBadRequest
		}
		s.Storage.AddGauge(server.Gauge{Metric: metric, Name: name, Value: valueFloat})
	}
	return nil

}

func (s *Service) ValueMetric(vars map[string]string) (string, error) {
	metric := vars["metric"]
	name := vars["name"]

	if metric == "counter" {
		value, err := s.Storage.GetCounter(name)
		if err != nil {
			return "", ErrStatusNotFound

		}
		return fmt.Sprint(value), nil
	}
	if metric == "gauge" {
		value, err := s.Storage.GetGauge(name)
		if err != nil {
			return "", ErrStatusNotFound

		}
		return fmt.Sprint(value), nil
	}
	return "", ErrStatusNotFound
}

func (s *Service) GetAllMetrics() string {
	counter, gouge := s.Storage.GetAllMetrics()
	var b strings.Builder

	b.WriteString("counter")
	b.WriteByte('\n')
	b.WriteString("____________")
	b.WriteByte('\n')
	for k, v := range counter {
		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(fmt.Sprint(v))
		b.WriteByte('\n')
	}
	b.WriteString("gouge")
	b.WriteByte('\n')
	b.WriteString("____________")
	b.WriteByte('\n')
	for k, v := range gouge {
		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(fmt.Sprint(v))
		b.WriteByte('\n')
	}
	metrics := b.String()
	return metrics

}
