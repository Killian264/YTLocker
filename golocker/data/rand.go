package data

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strconv"
)

type DataRand interface {
	ID() uint64
}

// DataRand generates large random numbers
type ActualRand struct{}

func (r *ActualRand) ID() uint64 {
	buf := make([]byte, 8)

	rand.Read(buf)

	num := binary.LittleEndian.Uint64(buf)

	if(num > 99999999999999999){
		str := fmt.Sprint(num)[:15]
		num, _ = strconv.ParseUint(str, 10, 64)
	}

	return num
}

// DataRand generates large random numbers
type TestRand struct{}

// generates a number and cuts it to length 9
// 1 in 18 trillion chance the initial number is too small
func (r *TestRand) ID() uint64 {
	buf := make([]byte, 8)

	rand.Read(buf)

	num := binary.LittleEndian.Uint64(buf)

	str := fmt.Sprint(num)[:9]

	num, _ = strconv.ParseUint(str, 10, 64)

	return num
}
