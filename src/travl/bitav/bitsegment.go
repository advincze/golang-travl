package bitav

import (
	"bytes"
	"math/big"
	"strconv"
)

type BitSegment struct {
	big.Int
	start int
}

func (s *BitSegment) String() string {
	var buffer bytes.Buffer

	for i := 0; i < s.BitLen(); i++ {
		buffer.WriteString(strconv.Itoa(int(s.Bit(i))))
	}

	return buffer.String()
}

func NewBitSegment(start int) *BitSegment {
	return &BitSegment{Int: *big.NewInt(0), start: start}
}
