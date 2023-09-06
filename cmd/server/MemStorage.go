package main

type gauge gloat64
type counter int64

type MemStorage struct {
	st map[any]map[string]any
}
