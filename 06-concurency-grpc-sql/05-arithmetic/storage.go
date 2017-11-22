package main

import "fmt"

type storage struct {
	values []float64
}

func newStorage() *storage {
	return &storage{values: []float64{}}
}

func (s *storage) Get(pos int) (float64, error) {
	if len(s.values) < pos {
		return 0, fmt.Errorf("Index %d out of reach in storage", pos)
	}

	return s.values[pos-1], nil
}

func (s *storage) Append(val float64) int {
	s.values = append(s.values, val)

	return len(s.values)
}
