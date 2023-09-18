package services

import (
	"github.com/MorZLE/metrick/internal/server/storages"
	"testing"
)

func TestService_ProcessingMetrick(t *testing.T) {
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
					"value":  "3.2",
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
			name: "FailTest2_NotName",
			fields: fields{
				Storage: storages.NewStorage(),
			},
			args: args{
				vars: map[string]string{
					"metric": "counter",
					"name":   "",
					"value":  "3",
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
				Storage: tt.fields.Storage,
			}
			if err := s.ProcessingMetrick(tt.args.vars); (err != nil) != tt.wantErr {
				t.Errorf("ProcessingMetrick() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
