package storages

import (
	"github.com/MorZLE/metrick/internal/server/services"
)

func NewStorage() MemStorage {
	return MemStorage{make(map[string]float64)}
}

type MemStorage struct {
	m map[string]float64
}

func (s *MemStorage) AddCounter(v services.Counter) {
	_, ok := s.m[v.Name]
	if ok {
		s.m[v.Name] = v.Value + s.m[v.Name]
	} else {
		s.m[v.Name] = v.Value
	}
}
func (s *MemStorage) AddGauge(v services.Counter) {
	s.m[v.Name] = v.Value
}
