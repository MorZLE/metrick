package services

import (
	"errors"
	"fmt"
	"github.com/MorZLE/metrick/internal/constants"
	"github.com/MorZLE/metrick/internal/server"
	"github.com/MorZLE/metrick/internal/server/storages"
	"strconv"
	"strings"
)

func NewService(s storages.Repositories) Service {
	return Service{storage: s}
}

type ServiceInterface interface {
	ProcessingMetric(metric, name, val string) error
	ValueMetric(metric, name string) (string, error)
	GetAllMetrics() string
	ValueMetricJSON(metric string, name string) (constants.Metrics, error)
}

type Service struct {
	storage storages.Repositories
}

func (s *Service) ProcessingMetric(metric, name, value string) error {

	switch metric {
	case "":
		return constants.ErrStatusNotFound
	case "gauge":
		valueFloat, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return constants.ErrBadRequest
		}
		s.storage.AddGauge(server.Gauge{Metric: metric, Name: name, Value: valueFloat})
	case "counter":
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return constants.ErrBadRequest
		}
		s.storage.AddCounter(server.Counter{Metric: metric, Name: name, Value: valueInt})
	default:
		return constants.ErrBadRequest
	}

	return nil
}

func (s *Service) ValueMetric(metric, name string) (string, error) {

	if metric == "counter" {
		value, err := s.storage.GetCounter(name)
		if errors.Is(err, constants.ErrStatusNotFound) {
			return "", constants.ErrStatusNotFound

		}
		return fmt.Sprint(value), nil
	}
	if metric == "gauge" {
		value, err := s.storage.GetGauge(name)
		if errors.Is(err, constants.ErrStatusNotFound) {
			return "", constants.ErrStatusNotFound

		}
		return fmt.Sprint(value), nil
	}
	return "", constants.ErrStatusNotFound
}

func (s *Service) GetAllMetrics() string {
	counter, gouge := s.storage.GetMetrics()
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

func (s *Service) ValueMetricJSON(metric, name string) (constants.Metrics, error) {
	val, err := s.ValueMetric(metric, name)
	if errors.Is(err, constants.ErrStatusNotFound) {
		return constants.Metrics{}, constants.ErrStatusNotFound

	}
	if metric == "counter" {
		num, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return constants.Metrics{}, constants.ErrParseInt
		}
		return constants.Metrics{
			ID:    name,
			MType: "counter",
			Delta: &num,
		}, nil
	}
	if metric == "gauge" {
		num, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return constants.Metrics{}, constants.ErrParseFloat
		}
		return constants.Metrics{
			ID:    name,
			MType: "gauge",
			Value: &num,
		}, nil
	}
	return constants.Metrics{}, constants.ErrStatusNotFound
}
