package services

import (
	"fmt"
	"github.com/MorZLE/metrick/internal/client/handlers"
	"github.com/MorZLE/metrick/internal/client/storages"
	"strconv"
	"time"
)

func NewService(s storages.MetricInterface, h handlers.HandleRequest, pollinterval int, reportInterval int) Service {

	return Service{Metric: s, Handler: h, PollInterval: pollinterval, ReportInterval: reportInterval}
}

type ServiceInterface interface {
	UpClient()
	SendRequest()
}

type Service struct {
	Metric         storages.MetricInterface
	Handler        handlers.HandleRequest
	PollInterval   int
	ReportInterval int
}

func (s *Service) UpClient() {
	for {
		startTime := time.Now()
		for {
			if time.Now().Unix()-startTime.Unix() < int64(s.ReportInterval) {
				s.Metric.UpdateMetric()
				time.Sleep(time.Duration(s.PollInterval) * time.Second)

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
