package storages

import (
	"github.com/MorZLE/metrick/internal/server"
	"math/rand"
	"runtime"
)

var memStats runtime.MemStats

func NewStorage() Metric {
	return Metric{MGauge: make(map[string]any), MCounter: make(map[string]int)}
}

type Metric struct {
	PollCount server.Counter
	MGauge    map[string]any
	MCounter  map[string]int
}

func (m *Metric) UpdateMetric() *Metric {
	m.MCounter["PollCount"]++

	m.MGauge["Alloc"] = memStats.Alloc
	m.MGauge["BuckHashSys"] = memStats.BuckHashSys
	m.MGauge["Frees"] = memStats.Frees
	m.MGauge["GCCPUFraction"] = memStats.GCCPUFraction
	m.MGauge["GCSys"] = memStats.GCSys
	m.MGauge["HeapAlloc"] = memStats.HeapAlloc
	m.MGauge["HeapIdle"] = memStats.HeapIdle
	m.MGauge["HeapInuse"] = memStats.HeapInuse
	m.MGauge["HeapObjects"] = memStats.HeapObjects
	m.MGauge["HeapReleased"] = memStats.HeapReleased
	m.MGauge["HeapSys"] = memStats.HeapSys
	m.MGauge["LastGC"] = memStats.LastGC
	m.MGauge["Lookups"] = memStats.Lookups
	m.MGauge["MCacheInuse"] = memStats.MCacheInuse
	m.MGauge["MCacheSys"] = memStats.MCacheSys
	m.MGauge["MSpanInuse"] = memStats.MSpanInuse
	m.MGauge["MSpanSys"] = memStats.MSpanSys
	m.MGauge["Mallocs"] = memStats.Mallocs
	m.MGauge["NextGC"] = memStats.NextGC
	m.MGauge["NumForcedGC"] = memStats.NumForcedGC
	m.MGauge["NumGC"] = memStats.NumGC
	m.MGauge["OtherSys"] = memStats.OtherSys
	m.MGauge["PauseTotalNs"] = memStats.PauseTotalNs
	m.MGauge["StackInuse"] = memStats.StackInuse
	m.MGauge["StackSys"] = memStats.StackSys
	m.MGauge["Sys"] = memStats.Sys
	m.MGauge["TotalAlloc"] = memStats.TotalAlloc
	m.MGauge["RandomValue"] = rand.Float64()

	return m
}
