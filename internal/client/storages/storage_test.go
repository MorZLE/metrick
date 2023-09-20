package storages

import (
	"github.com/MorZLE/metrick/internal/server"
	"reflect"
	"testing"
)

func TestMetric_UpdateMetric(t *testing.T) {
	type fields struct {
		PollCount server.Counter
		MGauge    map[string]any
		MCounter  map[string]int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Metric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Metric{
				PollCount: tt.fields.PollCount,
				MGauge:    tt.fields.MGauge,
				MCounter:  tt.fields.MCounter,
			}
			if got := m.UpdateMetric(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateMetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStorage(t *testing.T) {
	tests := []struct {
		name string
		want Metric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStorage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}
