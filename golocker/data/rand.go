package data

import (
	"crypto/rand"
	"encoding/binary"
)

type DataRand interface {
	ID() uint64
}

type TestRand struct{ id uint64 }

func (r *TestRand) ID() uint64 {
	r.id++
	return r.id
}

type ActualRand struct{}

func (r *ActualRand) ID() uint64 {
	buf := make([]byte, 8)
	rand.Read(buf)
	return binary.LittleEndian.Uint64(buf)
}
