package server

import (
	"errors"
	"github.com/MorZLE/metrick/internal/storages/server"
)

func NewService(s server.Storage) Service {
	return Service{Storage: s}
}

type Gauge struct {
	metric string
	Name   string
	Value  float64
}

type Counter struct {
	metric string
	Name   string
	Value  int64
}

type Service struct {
	Storage server.Storage
}

func (s Service) VerifType(metric, name, value string) (struct{}, error) {
	if metric != "gauge" && metric != "counter" {
		return nil, errors.New("http.StatusBadRequest")
	}
	if metric != "gauge" {
		return Gauge{metric, name, int64(value)}, nil
	}
	if metric != "counter" {
		return Counter{metric, name, float64(value)}, nil
	}
	return nil, _
}

func (s Service) AddValue(metric, name, value string) error {
	t, err := s.VerifType(metric, name, value)
	if err != nil {
		return err
	}
	s.Storage.Add(t)
	return nil

}
