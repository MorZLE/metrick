package server

type Gauge struct {
	Metric string
	Name   string
	Value  float64
}

type Counter struct {
	Metric string
	Name   string
	Value  int
}
