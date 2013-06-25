package av

import (
	"bytes"
	"math/big"
	"strconv"
	"time"
)

type BitAv3 struct {
	internalRes TimeResolution
	segments    map[int]*Segment
}

func (av *BitAv3) size() int {
	var size int
	for _, segment := range av.segments {
		size += len(segment.Bytes())
	}
	return size
}

func (av *BitAv3) Set(from, to time.Time, value byte) {
	fromUnit := timeToUnit(from, av.internalRes)
	toUnit := timeToUnit(to, av.internalRes)
	av.SetAv(fromUnit, toUnit, value)
}

func (av *BitAv3) Get(from, to time.Time, res TimeResolution) *BitVector {
	fromUnit := timeToUnit(from, av.internalRes)
	toUnit := timeToUnit(to, av.internalRes)
	arr := av.GetAv(fromUnit, toUnit)
	bv := &BitVector{
		Resolution: res,
		Data:       arr,
		Start:      floorDate(from, res)}
	return bv
}

func (av *BitAv3) SetAt(at time.Time, value byte) {
	atUnit := timeToUnit(at, av.internalRes)
	av.SetAv(atUnit, atUnit+1, value)
}

func (av *BitAv3) GetAt(at time.Time) byte {
	atUnit := timeToUnit(at, av.internalRes)
	arr := av.GetAv(atUnit, atUnit+1)
	return byte(arr[0])
}

const segmentSize = 60 * 24

type Segment struct {
	big.Int
	start int
}

func (av *BitAv3) String() string {
	var buffer bytes.Buffer

	for _, segment := range av.segments {
		buffer.WriteString(strconv.Itoa(segment.start))
		buffer.WriteString("->")
		buffer.WriteString(segment.String())
		buffer.WriteRune('\n')
	}

	return buffer.String()
}

func (s *Segment) String() string {
	var buffer bytes.Buffer

	for i := 0; i < s.BitLen(); i++ {
		buffer.WriteString(strconv.Itoa(int(s.Bit(i))))
	}

	return buffer.String()
}

func NewSegment(start int) *Segment {
	return &Segment{Int: *big.NewInt(0), start: start}
}

func NewBitAv3() *BitAv3 {
	return &BitAv3{
		segments:    make(map[int]*Segment),
		internalRes: Minute,
	}
}

func segmentStart(i int) int {
	return i - i%segmentSize
}

func (av *BitAv3) getOrCreateSegment(startValue int) *Segment {
	if segment, ok := av.segments[startValue]; ok {
		return segment
	} else {
		segment = NewSegment(startValue)
		av.segments[startValue] = segment
		return segment
	}
}

func (av *BitAv3) SetAv(from, to int, value byte) {
	currentSegment := av.getOrCreateSegment(segmentStart(from))
	for i, j := from, from%segmentSize; i < to; i, j = i+1, j+1 {
		if j == segmentSize {
			currentSegment = av.getOrCreateSegment(i)
			j = 0
		}
		currentSegment.SetBit(&currentSegment.Int, j, uint(value))
	}

}

func (av *BitAv3) GetAv(from, to int) []byte {
	length := to - from
	result := make([]byte, length)
	currentSegment := av.getOrCreateSegment(segmentStart(from))
	for i, j := 0, from%segmentSize; i < length; i, j = i+1, j+1 {
		if j == segmentSize {
			currentSegment = av.getOrCreateSegment(i + from)
			j = 0
		}
		result[i] = byte(currentSegment.Bit(j))
	}
	return result
}
