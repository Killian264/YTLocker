package data

import (
	"crypto/rand"
	"encoding/binary"
)

type DataRand interface {
	ID() uint64
}

// TestRand generates non random small numbers for sqlite usage
type TestRand struct{ id uint64 }

func (r *TestRand) ID() uint64 {
	r.id++
	return r.id
}

// DataRand generates large random numbers
type ActualRand struct{}

func (r *ActualRand) ID() uint64 {
	buf := make([]byte, 8)
	rand.Read(buf)
	return binary.LittleEndian.Uint64(buf)
}
