package handlers

import (
	"github.com/MorZLE/metrick/internal/server/services"
	"net/http"
	"reflect"
	"testing"
)

func TestHandler_UpServer(t *testing.T) {
	type fields struct {
		Logic services.Service
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{
				Logic: tt.fields.Logic,
			}
			h.UpServer()
		})
	}
}

func TestHandler_UpdatePage(t *testing.T) {
	type fields struct {
		Logic services.Service
	}
	type args struct {
		res http.ResponseWriter
		req *http.Request
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
			h := Handler{
				Logic: tt.fields.Logic,
			}
			h.UpdatePage(tt.args.res, tt.args.req)
		})
	}
}

func TestHandler_routs(t *testing.T) {
	type fields struct {
		Logic services.Service
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{
				Logic: tt.fields.Logic,
			}
			h.routs()
		})
	}
}

func TestNewHandler(t *testing.T) {
	type args struct {
		l services.Service
	}
	tests := []struct {
		name string
		args args
		want Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
