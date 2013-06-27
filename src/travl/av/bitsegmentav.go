package av

import (
	"bytes"
	"strconv"
	"time"
)

type BitSegmentAv struct {
	ID          string
	internalRes TimeResolution
	segments    map[int]*BitSegment
	segmentSize int
}

func NewBitSegmentAv(ID string, res TimeResolution) *BitSegmentAv {
	return &BitSegmentAv{
		ID:          ID,
		segments:    make(map[int]*BitSegment),
		internalRes: res,
		segmentSize: int(Day / res),
	}
}

func (av *BitSegmentAv) size() int {
	var size int
	for _, segment := range av.segments {
		size += len(segment.Bytes())
	}
	return size
}

func (av *BitSegmentAv) Set(from, to time.Time, value byte) {
	fromUnit := timeToUnitFloor(from, av.internalRes)
	toUnit := timeToUnitFloor(to, av.internalRes)
	av.setUnitInternal(fromUnit, toUnit, value)
	//println("size after set:", av.size())
}

func (av *BitSegmentAv) Get(from, to time.Time, res TimeResolution) *BitVector {

	if res > av.internalRes {
		//15min > 5min
		// lower resolution
		// println("get w lower res")

		fromUnit := timeToUnitFloor(floorDate(from, res), av.internalRes)
		toUnit := timeToUnitFloor(ceilDate(to, res), av.internalRes)
		// println("from", from.String(), fromUnit)
		// println("to", to.String(), toUnit)
		arr := av.getUnitInternal(fromUnit, toUnit)
		// println("arr", len(arr))

		// internal:[0,0,0,1,1,1,0,0,0]
		// -> res:  [--0--,--1--,--0--]
		factor := int(res / av.internalRes)
		// println("factor", factor)
		reducedArr := reduceByFactor(arr, factor, reduceAllOne)
		bv := &BitVector{
			Resolution: res,
			Data:       reducedArr,
			Start:      floorDate(from, res)}
		return bv
	} else if res < av.internalRes {
		// 1min < 5min
		// higher resolution
		// println("get w higher res", from.String(), to.String())
		fromUnitInternalRes := timeToUnitFloor(from, av.internalRes)
		toUnitInternalRes := timeToUnitFloor(ceilDate(to, av.internalRes), av.internalRes)
		// println("from, tounit ", fromUnitInternalRes, toUnitInternalRes)
		arr := av.getUnitInternal(fromUnitInternalRes, toUnitInternalRes)
		// println("arr: ", printarr(arr))
		// internal: [0        ,1          ,1        ]
		// -> res  : [0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1]
		//  multiply by factor
		factor := int(av.internalRes / res)
		// println("factor: ", factor)
		arrMultiplied := multiplyByFactor(arr, factor)
		// println("arr2: ", printarr(arrMultiplied))

		cutoff := timeToUnitFloor(from, res) - fromUnitInternalRes*factor
		// println("cutoff: ", cutoff)
		origlen := timeToUnitFloor(to, res) - timeToUnitFloor(from, res)
		// println("origlen: ", origlen)

		arrTrimmed := arrMultiplied[cutoff : cutoff+origlen]
		bv := &BitVector{
			Resolution: res,
			Data:       arrTrimmed,
			Start:      floorDate(from, av.internalRes)}
		return bv
		// TODO cut off excess:
		//   [0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1]
		// ->      [0,0,1,1,1,1,1,1,1,1,1]
	} else {
		// internal resolution
		// println("get w internal res")
		fromUnit := timeToUnitFloor(from, res)
		toUnit := timeToUnitFloor(to, res)
		arr := av.getUnitInternal(fromUnit, toUnit)
		bv := &BitVector{
			Resolution: res,
			Data:       arr,
			Start:      floorDate(from, res)}
		return bv
	}
}

func printarr(arr []byte) string {
	var buffer bytes.Buffer

	buffer.WriteString("[")
	buffer.WriteString(strconv.Itoa(len(arr)))
	buffer.WriteString("-")
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			buffer.WriteString("1, ")
		} else {
			buffer.WriteString("0, ")
		}
	}
	buffer.WriteString("],")

	return buffer.String()
}

func multiplyByFactor(data []byte, factor int) []byte {
	length := len(data) * factor
	var multipliedData []byte = make([]byte, length)
	j := 0
	for _, b := range data {

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

func reduceAnyOne(data []byte) byte {
	for _, b := range data {
		if b == 1 {
			return 1
		}
	}
	return 0
}

func reduceMajority(data []byte) byte {
	sizewin := len(data) / 2
	count := 0
	for _, b := range data {
		if b == 1 {
			count++
		}
	}
	if count > sizewin {
		return 1
	}
	return 0
}

func (av *BitSegmentAv) SetAt(at time.Time, value byte) {
	atUnit := timeToUnitFloor(at, av.internalRes)
	av.setUnitInternal(atUnit, atUnit+1, value)
}

func (av *BitSegmentAv) GetAt(at time.Time) byte {
	atUnit := timeToUnitFloor(at, av.internalRes)
	arr := av.getUnitInternal(atUnit, atUnit+1)
	return byte(arr[0])
}

func (av *BitSegmentAv) String() string {
	var buffer bytes.Buffer

	for _, segment := range av.segments {
		buffer.WriteString(strconv.Itoa(segment.start))
		buffer.WriteString("->")
		buffer.WriteString(segment.String())
		buffer.WriteRune('\n')
	}

	return buffer.String()
}

func (av *BitSegmentAv) segmentStart(i int) int {
	return i - i%av.segmentSize
}

func (av *BitSegmentAv) getOrCreateBitSegment(startValue int) *BitSegment {
	if segment, ok := av.segments[startValue]; ok {
		return segment
	} else {
		segment = NewBitSegment(startValue)
		av.segments[startValue] = segment
		return segment
	}
}

func (av *BitSegmentAv) getOrEmptyBitSegment(startValue int) *BitSegment {
	if segment, ok := av.segments[startValue]; ok {
		return segment
	} else {
		return NewBitSegment(startValue)
	}
}

func (av *BitSegmentAv) setUnitInternal(from, to int, value byte) {
	currentBitSegment := av.getOrCreateBitSegment(av.segmentStart(from))
	for i, j := from, from%av.segmentSize; i < to; i, j = i+1, j+1 {
		if j == av.segmentSize {
			currentBitSegment = av.getOrCreateBitSegment(i)
			j = 0
		}
		currentBitSegment.SetBit(&currentBitSegment.Int, j, uint(value))
	}
	//println(av.String())
}

func (av *BitSegmentAv) getUnitInternal(from, to int) []byte {
	length := to - from
	result := make([]byte, length)
	currentBitSegment := av.getOrEmptyBitSegment(av.segmentStart(from))
	for i, j := 0, from%av.segmentSize; i < length; i, j = i+1, j+1 {
		if j == av.segmentSize {
			currentBitSegment = av.getOrEmptyBitSegment(i + from)
			j = 0
		}
		result[i] = byte(currentBitSegment.Bit(j))
	}
	return result
}
