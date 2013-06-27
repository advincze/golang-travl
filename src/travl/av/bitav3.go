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
	segmentSize int
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
	println("size after set:", av.size())
}

func (av *BitAv3) Get(from, to time.Time, res TimeResolution) *BitVector {

	if res > av.internalRes {
		//15min > 5min
		// lower resolution
		println("get w lower res")

		fromUnit := timeToUnit(floorDate(from, res), av.internalRes)
		toUnit := timeToUnit(floorDate(to, res), av.internalRes)
		println("from", from.String(), fromUnit)
		println("to", to.String(), toUnit)
		arr := av.GetAv(fromUnit, toUnit)
		println("arr", len(arr))

		// internal:[0,0,0,1,1,1,0,0,0]
		// -> res:  [--0--,--1--,--0--]
		factor := int(res / av.internalRes)
		println("factor", factor)
		reducedArr := reduceByFactor(arr, factor, reduceAllOne)
		bv := &BitVector{
			Resolution: res,
			Data:       reducedArr,
			Start:      floorDate(from, res)}
		return bv
	} else if res < av.internalRes {
		// 1min < 5min
		// higher resolution
		fromUnit := timeToUnit(from, av.internalRes)
		toUnit := timeToUnit(to, av.internalRes)
		arr := av.GetAv(fromUnit, toUnit)
		// internal: [0        ,1          ,1        ]
		// -> res  : [0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1]
		//  multiply by factor
		factor := int(av.internalRes / res)
		arrMultiplied := multiplyByFactor(arr, factor)
		bv := &BitVector{
			Resolution: res,
			Data:       arrMultiplied,
			Start:      floorDate(from, av.internalRes)}
		return bv
		// TODO cut off excess:
		//   [0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1]
		// ->      [0,0,1,1,1,1,1,1,1,1,1]
	} else {
		// internal resolution
		fromUnit := timeToUnit(from, av.internalRes)
		toUnit := timeToUnit(to, av.internalRes)
		arr := av.GetAv(fromUnit, toUnit)
		bv := &BitVector{
			Resolution: res,
			Data:       arr,
			Start:      floorDate(from, res)}
		return bv
	}
}

func multiplyByFactor(data []byte, factor int) []byte {
	length := len(data) * factor
	var multipliedData []byte = make([]byte, length)
	for _, b := range data {
		j := 0
		for i := 0; i < factor; i++ {
			multipliedData[j] = b
			j++
		}
	}
	return multipliedData
}

func reduceByFactor(data []byte, factor int, reduceFn func([]byte) byte) []byte {
	length := len(data) / factor
	var reducedData []byte = make([]byte, length)
	for i, j := 0, 0; i < length; i++ {
		reducedData[i] = reduceFn(data[j : j+factor])
		j += factor
	}
	return reducedData
}

func reduceAllOne(data []byte) byte {
	for _, b := range data {
		if b != 1 {
			return 0
		}
	}
	return 1
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

//const segmentSize = 60 * 24

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

func NewBitAv3(res TimeResolution) *BitAv3 {
	return &BitAv3{
		segments:    make(map[int]*Segment),
		internalRes: res,
		segmentSize: 100,
	}
}

func (av *BitAv3) segmentStart(i int) int {
	return i - i%av.segmentSize
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
	currentSegment := av.getOrCreateSegment(av.segmentStart(from))
	for i, j := from, from%av.segmentSize; i < to; i, j = i+1, j+1 {
		if j == av.segmentSize {
			currentSegment = av.getOrCreateSegment(i)
			j = 0
		}
		currentSegment.SetBit(&currentSegment.Int, j, uint(value))
	}

}

func (av *BitAv3) GetAv(from, to int) []byte {
	length := to - from
	result := make([]byte, length)
	currentSegment := av.getOrCreateSegment(av.segmentStart(from))
	for i, j := 0, from%av.segmentSize; i < length; i, j = i+1, j+1 {
		if j == av.segmentSize {
			currentSegment = av.getOrCreateSegment(i + from)
			j = 0
		}
		result[i] = byte(currentSegment.Bit(j))
	}
	return result
}
