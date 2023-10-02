package storages

import (
	"math/rand"
	"runtime"
)

var memStats runtime.MemStats

func NewStorage() Metric {
	return Metric{mGauge: make(map[string]any), mCounter: make(map[string]int)}
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=MetricInterface
type MetricInterface interface {
	UpdateMetric() *Metric
	GetMGauge() map[string]any
	GetMCounter() map[string]int
}

type Metric struct {
	mGauge   map[string]any
	mCounter map[string]int
}

func (m *Metric) UpdateMetric() *Metric {
	m.mCounter["PollCount"]++

	m.mGauge["Alloc"] = memStats.Alloc
	m.mGauge["BuckHashSys"] = memStats.BuckHashSys
	m.mGauge["Frees"] = memStats.Frees
	m.mGauge["GCCPUFraction"] = memStats.GCCPUFraction
	m.mGauge["GCSys"] = memStats.GCSys
	m.mGauge["HeapAlloc"] = memStats.HeapAlloc
	m.mGauge["HeapIdle"] = memStats.HeapIdle
	m.mGauge["HeapInuse"] = memStats.HeapInuse
	m.mGauge["HeapObjects"] = memStats.HeapObjects
	m.mGauge["HeapReleased"] = memStats.HeapReleased
	m.mGauge["HeapSys"] = memStats.HeapSys
	m.mGauge["LastGC"] = memStats.LastGC
	m.mGauge["Lookups"] = memStats.Lookups
	m.mGauge["MCacheInuse"] = memStats.MCacheInuse
	m.mGauge["MCacheSys"] = memStats.MCacheSys
	m.mGauge["MSpanInuse"] = memStats.MSpanInuse
	m.mGauge["MSpanSys"] = memStats.MSpanSys
	m.mGauge["Mallocs"] = memStats.Mallocs
	m.mGauge["NextGC"] = memStats.NextGC
	m.mGauge["NumForcedGC"] = memStats.NumForcedGC
	m.mGauge["NumGC"] = memStats.NumGC
	m.mGauge["OtherSys"] = memStats.OtherSys
	m.mGauge["PauseTotalNs"] = memStats.PauseTotalNs
	m.mGauge["StackInuse"] = memStats.StackInuse
	m.mGauge["StackSys"] = memStats.StackSys
	m.mGauge["Sys"] = memStats.Sys
	m.mGauge["TotalAlloc"] = memStats.TotalAlloc
	m.mGauge["RandomValue"] = rand.Float64()

	return m
}

func (m *Metric) GetMGauge() map[string]any {
	return m.mGauge
}

func (m Metric) GetMCounter() map[string]int {
	return m.mCounter
}
