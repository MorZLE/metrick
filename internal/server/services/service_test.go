package services

import (
	"github.com/MorZLE/metrick/internal/constants"
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
				storage: &tt.fields.Storage,
			}
			if err := s.ProcessingMetric(tt.args.vars["metric"], tt.args.vars["name"], tt.args.vars["value"]); (err != nil) != tt.wantErr {
				t.Errorf("ProcessingMetrick() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_ValueMetric(t *testing.T) {

	type mckS func(r *mocks.Repositories)

	type args struct {
		vars        map[string]string
		mockStorage mckS
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
				},
				mockStorage: func(r *mocks.Repositories) {
					r.On("GetCounter", "test1").Return(23, nil).Once()
				},
			},

			want:    "23",
			wantErr: false,
		},

		{
			name: "PositiveTest2",
			args: args{
				vars: map[string]string{
					"metric": "gauge",
					"name":   "test2",
				},
				mockStorage: func(r *mocks.Repositories) {
					r.On("GetGauge", "test2").Return(23.3, nil).Once()
				},
			},
			want:    "23.3",
			wantErr: false,
		},
		{
			name: "FailTestGetCounter",
			args: args{
				vars: map[string]string{
					"metric": "counter",
					"name":   "test3",
				},
				mockStorage: func(r *mocks.Repositories) {
					r.On("GetCounter", "test3").Return(0, constants.ErrStatusNotFound).Once()
				},
			},
			wantErr: true,
		},

		{
			name: "FailTestGetGauge",
			args: args{
				vars: map[string]string{
					"metric": "gauge",
					"name":   "test4",
				},
				mockStorage: func(r *mocks.Repositories) {
					r.On("GetGauge", "test4").Return(0.0, constants.ErrStatusNotFound).Once()
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := mocks.NewRepositories(t)
			tt.args.mockStorage(storage)

			s := &Service{
				storage: storage,
			}

			got, err := s.ValueMetric(tt.args.vars["metric"], tt.args.vars["name"])

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
