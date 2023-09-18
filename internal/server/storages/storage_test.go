package storages

import (
	"github.com/MorZLE/metrick/internal/server"
	"testing"
)

func TestMemStorage_AddCounter(t *testing.T) {
	type fields struct {
		mCounter map[string]int
		mGouge   map[string]float64
	}
	type args struct {
		v server.Counter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MemStorage{
				mCounter: tt.fields.mCounter,
				mGouge:   tt.fields.mGouge,
			}
			s.AddCounter(tt.args.v)
		})
	}
}

func TestMemStorage_AddGauge(t *testing.T) {
	type fields struct {
		mCounter map[string]int
		mGouge   map[string]float64
	}
	type args struct {
		v server.Gauge
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MemStorage{
				mCounter: tt.fields.mCounter,
				mGouge:   tt.fields.mGouge,
			}
			s.AddGauge(tt.args.v)
		})
	}
}