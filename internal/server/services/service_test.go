package services

import (
	"github.com/MorZLE/metrick/internal/server/storages"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		s storages.MemStorage
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		// TODO: Add test cases.
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
