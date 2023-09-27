package services

import (
	"github.com/MorZLE/metrick/config"
	mocks2 "github.com/MorZLE/metrick/internal/client/mocks"
	"testing"
)

func TestService_SendRequest(t *testing.T) {

	type args struct {
		metric string
		name   string
		val    string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "goodTest1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			client := mocks2.NewHandleRequest(t)
			storage := mocks2.NewMetricInterface(t)

			storage.On("GetMCounter").Return(map[string]int{
				"erg": 23,
			}).Once()
			storage.On("GetMGauge").Return(map[string]interface{}{
				"wer": 23.3,
			}).Once()

			s := &Service{
				Handler: client,
				Metric:  storage,
				cnf: &config.ConfigAgent{
					FlagAddr: ":8080",
				},
			}

			client.On("Request", "gauge", "wer", "23.3", ":8080").Return().Once()
			client.On("Request", "counter", "erg", "23", ":8080").Return().Once()

			s.SendRequest()
		})
	}
}
