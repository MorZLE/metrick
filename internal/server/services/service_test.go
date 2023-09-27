package services

import (
	"errors"
	"github.com/MorZLE/metrick/internal/server/mocks"
	"github.com/MorZLE/metrick/internal/server/storages"
	"testing"
)

func TestService_ProcessingMetric(t *testing.T) {
	type fields struct {
		Storage storages.MemStorage
	}
	type args struct {
		vars map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "PositiveTest1",
			fields: fields{
				Storage: storages.NewStorage(),
			},
			args: args{
				vars: map[string]string{
					"metric": "gauge",
					"name":   "test1",
					"value":  "3.18675",
				}},

			wantErr: false,
		},
		{
			name: "PositiveTest2",
			fields: fields{
				Storage: storages.NewStorage(),
			},
			args: args{
				vars: map[string]string{
					"metric": "counter",
					"name":   "test2",
					"value":  "3",
				}},

			wantErr: false,
		},
		{
			name: " FailTest1_TypeMetric",
			fields: fields{
				Storage: storages.NewStorage(),
			},
			args: args{
				vars: map[string]string{
					"metric": "Gof",
					"name":   "test3",
					"value":  "3.2",
				}},

			wantErr: true,
		},
		{
			name: "FailTest3_NotName",
			fields: fields{
				Storage: storages.NewStorage(),
			},
			args: args{
				vars: map[string]string{
					"metric": "counter",
					"name":   "test3",
					"value":  "awd",
				}},

			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				Storage: &tt.fields.Storage,
			}
			if err := s.ProcessingMetric(tt.args.vars); (err != nil) != tt.wantErr {
				t.Errorf("ProcessingMetrick() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_ValueMetric(t *testing.T) {

	type args struct {
		vars map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "PositiveTest1",
			args: args{
				vars: map[string]string{
					"metric": "counter",
					"name":   "test1",
				}},
			want:    "23",
			wantErr: false,
		},

		{
			name: "PositiveTest1",
			args: args{
				vars: map[string]string{
					"metric": "gauge",
					"name":   "test2",
				}},
			want:    "23.3",
			wantErr: false,
		},
		{
			name: "FailTest1",
			args: args{
				vars: map[string]string{
					"metric": "counter",
					"name":   "test3",
				}},
			wantErr: true,
		},

		{
			name: "FailTest2",
			args: args{
				vars: map[string]string{
					"metric": "gauge",
					"name":   "test4",
				}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := mocks.NewRepositories(t)
			switch tt.args.vars["metric"] {
			case "counter":

				if tt.wantErr {
					storage.On("GetCounter", tt.args.vars["name"]).Return(0, errors.New("not found")).Once()
				} else {
					storage.On("GetCounter", tt.args.vars["name"]).Return(23, nil).Once()
				}
			case "gauge":

				if tt.wantErr {

					storage.On("GetGauge", tt.args.vars["name"]).Return(0.0, errors.New("not found")).Once()
				} else {
					storage.On("GetGauge", tt.args.vars["name"]).Return(23.3, nil).Once()
				}

			}

			s := &Service{
				Storage: storage,
			}

			got, err := s.ValueMetric(tt.args.vars)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValueMetric() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValueMetric() got = %v, want %v", got, tt.want)
			}

		})
	}
}
