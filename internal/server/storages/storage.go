package storages

import "github.com/MorZLE/metrick/internal/server"

func NewStorage() MemStorage {
	mC := make(map[string]int)
	mG := make(map[string]float64)
	return MemStorage{mCounter: mC, mGouge: mG}
}

//type repositories interface {
//	AddCounter()
//	AddGauge()
//}

type MemStorage struct {
	mCounter map[string]int
	mGouge   map[string]float64
}

func (s *MemStorage) AddCounter(v server.Counter) {
	s.mCounter[v.Name] += v.Value

}

func (s *MemStorage) AddGauge(v server.Gauge) {
	s.mGouge[v.Name] = v.Value
}
