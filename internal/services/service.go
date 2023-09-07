package services

import (
	"github.com/MorZLE/metrick/internal/storages"
)

type Service struct {
	metric  string
	name    string
	value   string
	Storage storages.Storage
}

func (s Service) New(rep storages.Storage) *Service {
	return &Service{Storage: rep}
}
