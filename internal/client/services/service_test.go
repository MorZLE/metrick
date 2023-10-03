package services

import (
	"github.com/MorZLE/metrick/config"
	"github.com/MorZLE/metrick/internal/client/constants"
	"github.com/MorZLE/metrick/internal/client/mocks"
	"testing"
)

func TestService_SendRequest(t *testing.T) {
	t1 := 23.3
	t2 := 234.34234

	var t11 int64 = 23
	var t22 int64 = 2346436436

	type mckS func(r *mocks.MetricInterface)
	type mckH func(r *mocks.HandleRequest)

	type args struct {
		mockHandler mckH
		mockStorage mckS
	}

	type metArgs constants.Metrics

	obj := map[string]metArgs{
		"test1": {
			ID:    "wer",
			MType: "gauge",
			Delta: nil,
			Value: &t1,
		},
		"test1.1": {
			ID:    "erg",
			MType: "counter",
			Delta: &t11,
			Value: nil,
		},
		"test2": {
			ID:    "sdfwefvdv",
			MType: "gauge",
			Delta: nil,
			Value: &t2,
		},
		"test2.1": {
			ID:    "segfrdbhtfhtrh",
			MType: "counter",
			Delta: &t22,
			Value: nil,
		},
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
					r.On("Request", obj["test1"], ":8080").Return().Once()
					r.On("Request", obj["test1.1"], ":8080").Return().Once()
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
					r.On("Request", obj["test2"], ":8080").Return().Once()
					r.On("Request", obj["test2.1"], ":8080").Return().Once()
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.NewHandleRequest(t)
			storage := mocks.NewMetricInterface(t)

			tt.args.mockStorage(storage)
			tt.args.mockHandler(client)

			s := &Service{
				handler: client,
				metric:  storage,
				cnf: &config.ConfigAgent{
					FlagAddr: ":8080",
				},
			}

			s.SendRequest()
		})
	}
}
