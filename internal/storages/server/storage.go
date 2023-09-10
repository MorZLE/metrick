package server

import (
	"github.com/MorZLE/metrick/internal/services/server"
)

type gauge float64
type counter int64

func NewStorage() Storage {
	return Storage{make(map[any]map[string]any)}
}

type Storage struct {
	st map[any]map[string]any
}

func (s Storage) Add(v server.Counter) {
	if v.Metric == "counter" {
		s.st[v.Metric][v.Name] = v.Value
	}
	if v.Metric == "gauge" {
		s.st[v.Metric][v.Name] += v.Value
	}
}
