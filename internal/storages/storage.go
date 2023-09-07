package storages

type gauge float64
type counter int64

type Storage struct {
	st map[any]map[string]any
}

func (s Storage) New() *Storage {
	return &Storage{}
}
