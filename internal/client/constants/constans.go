package constants

import "errors"

var ErrBadRequest = errors.New("BadRequest")
var ErrStatusNotFound = errors.New("StatusNotFound")
var ErrParseInt = errors.New("ErrParseInt")
var ErrParseFloat = errors.New("ErrParseFloat")

type Metrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}
