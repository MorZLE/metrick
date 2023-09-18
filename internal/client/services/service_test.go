package services

import (
	"github.com/MorZLE/metrick/internal/client/handlers"
	"github.com/MorZLE/metrick/internal/client/handlers/mocks"
	"github.com/MorZLE/metrick/internal/client/storages"
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
			args: args{
				"counter",
				"test",
				"1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mocks.NewHandleRequest(t)
			storage := mocks.NewMetricInterface(t)
			client.Request("counter", "test", "1")

			s := &Service{
				Handler: client,
				Metric:  storage,
			}
			s.SendRequest()
		})
	}
}

func TestService_UpClient(t *testing.T) {
	type fields struct {
		Metric  storages.Metric
		Handler handlers.Handler
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Metric:  tt.fields.Metric,
				Handler: tt.fields.Handler,
			}
			s.UpClient()
		})
	}
}
