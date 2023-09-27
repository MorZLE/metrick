package storages

import (
	"errors"
	"github.com/MorZLE/metrick/internal/server"
)

func NewStorage() MemStorage {
	mC := make(map[string]int)
	mG := make(map[string]float64)
	return MemStorage{mCounter: mC, mGouge: mG}
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=Repositories
type Repositories interface {
	AddCounter(v server.Counter)
	AddGauge(v server.Gauge)
	GetCounter(name string) (int, error)
	GetGauge(name string) (float64, error)

	GetMetrics() (map[string]int, map[string]float64)
}

type MemStorage struct {
	Repositories
	mCounter map[string]int
	mGouge   map[string]float64
}

func (s *MemStorage) AddCounter(v server.Counter) {
	s.mCounter[v.Name] += v.Value

}

func (s *MemStorage) AddGauge(v server.Gauge) {
	s.mGouge[v.Name] = v.Value
}

func (s *MemStorage) GetCounter(name string) (int, error) {
	if v, err := s.mCounter[name]; err {
		return v, nil
	}
	return 0, errors.New("not found")
}

func (s *MemStorage) GetGauge(name string) (float64, error) {
	if v, err := s.mGouge[name]; err {
		return v, nil
	}
	return 0.0, errors.New("not found")
}

func (s *MemStorage) GetMetrics() (map[string]int, map[string]float64) {
	return s.mCounter, s.mGouge

}
