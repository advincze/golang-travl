package bitav

import (
	"bytes"
	"math/big"
	"strconv"
)

type BitSegment struct {
	big.Int
	start int
	ID    string
}

func (s *BitSegment) String() string {
	var buffer bytes.Buffer

	for i := 0; i < s.BitLen(); i++ {
		buffer.WriteString(strconv.Itoa(int(s.Bit(i))))
	}

	return buffer.String()
}

func NewBitSegment(id string, start int) *BitSegment {
	return &BitSegment{
		Int:   *big.NewInt(0),
		start: start,
		ID:    id,
	}
}

var bitAvSegments map[string]map[int]*BitSegment

func (s *BitSegment) Save() {
	if bitAvSegments == nil {
		bitAvSegments = make(map[string]map[int]*BitSegment)
	}
	segments, ok := bitAvSegments[s.ID]
	if !ok {
		segments = make(map[int]*BitSegment)
		bitAvSegments[s.ID] = segments
	}
	segments[s.start] = s

}

func FindBitSegment(id string, start int) *BitSegment {
	if bitAvSegments == nil {
		return nil
	}
	segments, ok := bitAvSegments[id]
	if !ok {
		return nil
	}
	return segments[start]
}

func FindAllSegments(id string) map[int]*BitSegment {
	if bitAvSegments == nil {
		return nil
	}
	return bitAvSegments[id]
}
