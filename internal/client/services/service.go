package services

import (
	"fmt"
	"github.com/MorZLE/metrick/internal/client/handlers"
	"github.com/MorZLE/metrick/internal/client/storages"
	"strconv"
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
		startTime := time.Now()
		for {
			if time.Now().Unix()-startTime.Unix() < reportInterval {
				s.Metric.UpdateMetric()
				time.Sleep(pollInterval * time.Second)

			} else {
				break
			}
		}
		s.SendRequest()
	}
}

func (s *Service) SendRequest() {
	mGouge := s.Metric.GetMGauge()
	for k, v := range mGouge {
		s.Handler.Request("gauge", k, fmt.Sprint(v))
	}
	mCounter := s.Metric.GetMCounter()
	for k, v := range mCounter {
		s.Handler.Request("counter", k, strconv.Itoa(v))
	}

}
