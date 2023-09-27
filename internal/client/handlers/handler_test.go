package handlers

import (
	"testing"
)

func TestHandler_Request(t *testing.T) {
	type args struct {
		metric string
		name   string
		val    string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{}
			h.Request(tt.args.metric, tt.args.name, tt.args.val)
		})
	}
}
