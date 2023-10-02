package services

import (
	"github.com/MorZLE/metrick/config"
	"github.com/MorZLE/metrick/internal/client/mocks"
	"testing"
)

func TestService_SendRequest(t *testing.T) {

	type mckS func(r *mocks.MetricInterface)
	type mckH func(r *mocks.HandleRequest)

	type args struct {
		mockHandler mckH
		mockStorage mckS
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "goodTest1",
			args: args{
				mockStorage: func(r *mocks.MetricInterface) {
					r.On("GetMGauge").Return(map[string]interface{}{
						"wer": 23.3,
					}).Once()
					r.On("GetMCounter").Return(map[string]int{
						"erg": 23,
					}).Once()
				},
				mockHandler: func(r *mocks.HandleRequest) {
					r.On("Request", "gauge", "wer", "23.3", ":8080").Return().Once()
					r.On("Request", "counter", "erg", "23", ":8080").Return().Once()
				},
			},
		},
		{
			name: "goodTest2",
			args: args{
				mockStorage: func(r *mocks.MetricInterface) {
					r.On("GetMGauge").Return(map[string]interface{}{
						"sdfwefvdv": 234.34234,
					}).Once()
					r.On("GetMCounter").Return(map[string]int{
						"segfrdbhtfhtrh": 2346436436,
					}).Once()
				},
				mockHandler: func(r *mocks.HandleRequest) {
					r.On("Request", "gauge", "sdfwefvdv", "234.34234", ":8080").Return().Once()
					r.On("Request", "counter", "segfrdbhtfhtrh", "2346436436", ":8080").Return().Once()
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.NewHandleRequest(t)
			storage := mocks.NewMetricInterface(t)

			tt.args.mockStorage(storage)

			s := &Service{
				handler: client,
				metric:  storage,
				cnf: &config.ConfigAgent{
					FlagAddr: ":8080",
				},
			}

			tt.args.mockHandler(client)

			s.SendRequest()
		})
	}
}
