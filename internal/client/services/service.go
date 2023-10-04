package services

import (
	"fmt"
	"github.com/MorZLE/metrick/config"
	"github.com/MorZLE/metrick/internal/client/constants"
	"github.com/MorZLE/metrick/internal/client/handlers"
	"github.com/MorZLE/metrick/internal/client/storages"
	"log"
	"strconv"
	"time"
)

func NewService(s storages.MetricInterface, h handlers.HandleRequest, cnf *config.ConfigAgent) Service {

	return Service{metric: s, handler: h, cnf: cnf}
}

type ServiceInterface interface {
	UpClient()
	SendRequest()
}

type Service struct {
	metric  storages.MetricInterface
	handler handlers.HandleRequest
	cnf     *config.ConfigAgent
}

func (s *Service) UpClient() {
	for {
		startTime := time.Now()
		for {
			if time.Now().Unix()-startTime.Unix() < int64(s.cnf.FlagReportInterval) {
				s.metric.UpdateMetric()
				time.Sleep(time.Duration(s.cnf.FlagPollInterval) * time.Second)

			} else {
				break
			}
		}
		s.SendRequest()
	}
}

func (s *Service) SendRequest() {
	mGouge := s.metric.GetMGauge()
	for k, v := range mGouge {
		num, err := strconv.ParseFloat(fmt.Sprint(v), 64)
		if err != nil {
			log.Println(err)
		}
		req := constants.Metrics{ID: k, MType: "gauge", Value: &num}
		s.handler.Request(req, s.cnf.FlagAddr)

	}
	mCounter := s.metric.GetMCounter()
	for k, v := range mCounter {
		num := int64(v)
		req := constants.Metrics{ID: k, MType: "counter", Delta: &num}
		s.handler.Request(req, s.cnf.FlagAddr)
	}

}
