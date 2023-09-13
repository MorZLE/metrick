package storages

import (
	"github.com/MorZLE/metrick/internal/server"
	"math/rand"
	"runtime"
)

var memStats runtime.MemStats

func NewStorage() Metric {
	met := make(map[string]float64{
		"Alloc":         0.0,
		"BuckHashSys":   0.0,
		"Frees":         0.0,
		"GCCPUFraction": 0.0,
		"GCSys":         0.0,
		"HeapAlloc":     0.0,
		"HeapIdle":      0.0,
		"HeapInuse":     0.0,
		"HeapObjects":   0.0,
		"HeapReleased":  0.0,
		"HeapSys":       0.0,
		"LastGC":        0.0,
		"Lookups":       0.0,
		"MCacheInuse":   0.0,
		"MCacheSys":     0.0,
		"MSpanInuse":    0.0,
		"MSpanSys":      0.0,
		"Mallocs":       0.0,
		"NextGC":        0.0,
		"NumForcedGC":   0.0,
		"NumGC":         0.0,
		"OtherSys":      0.0,
		"PauseTotalNs":  0.0,
		"StackInuse":    0.0,
		"StackSys":      0.0,
		"Sys":           0.0,
		"TotalAlloc":    0.0,
	})
	pc := server.Counter{Metric: "Counter", Name: "PollCount", Value: 0}
	rv := server.Gauge{Metric: "Gauge", Name: "RandomValue"}
	return Metric{PollCount: pc, RandomValue: rv, Met: met}
}

type Metric struct {
	PollCount   server.Counter
	RandomValue server.Gauge
	Met         map[interface{}]server.Gauge
}

func (m *Metric) UpdateMetric() *Metric {
	m.PollCount.Value++
	m.RandomValue.Value = rand.Float64()
	for k := range m.Met {
		m.Met[k] = memStats.k
	}
	return m
}
