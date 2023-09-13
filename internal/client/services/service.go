package services

import (
	"github.com/MorZLE/metrick/internal/client/handlers"
	"github.com/MorZLE/metrick/internal/client/storages"
	"time"
)

func NewService(s storages.Metric, h handlers.Handler) Service {
	return Service{Metric: s, Handler: h}
}

type Service struct {
	Metric  storages.Metric
	Handler handlers.Handler
}

const pollInterval = 2
const reportInterval = 10

func (s *Service) UpClient() {
	for {
		time.Sleep(pollInterval * time.Second)
		s.Metric.UpdateMetric()
		time.Sleep(reportInterval * time.Second)
		s.SendRequest()
	}
}

func (s *Service) SendRequest() {
	for k, v := range s.Metric.Met {
		s.Handler.Request()

	}
	s.Handler.Request(s.Metric.RandomValue.Metric, s.Metric.RandomValue.Name, s.Metric.RandomValue.Value)
	s.Handler.Request(s.Metric.PollCount.Metric, s.Metric.PollCount.Name, s.Metric.PollCount.Value)
}
