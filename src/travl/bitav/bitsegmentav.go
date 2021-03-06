package bitav

import (
	"bytes"
	"strconv"
	"time"
	"travl/av"
)

type SegmentBitAv struct {
	ID          string
	internalRes av.TimeResolution
	segmentSize int
	bsp         BitSegmentPersistor
}

func NewSegmentBitAvDefault(id string) *SegmentBitAv {
	return NewSegmentBitAv(id, av.Minute5, "mem")
}

func NewSegmentBitAv(ID string, res av.TimeResolution, persistor string) *SegmentBitAv {
	var bsp BitSegmentPersistor
	switch persistor {
	case "mgo":
		bsp = new(BitSegmentMgoPersistor)
	default:
		bsp = new(BitSegmentMemPersistor)
	}
	return &SegmentBitAv{
		ID:          ID,
		internalRes: res,
		segmentSize: int(av.Day / res),
		bsp:         bsp,
	}
}

func (ba *SegmentBitAv) size() int {
	var size int
	segments := ba.bsp.FindAll(ba.ID)
	for _, segment := range segments {
		size += len(segment.Bytes())
	}
	return size
}

func (ba *SegmentBitAv) Set(from, to time.Time, value byte) {
	fromUnit := av.TimeToUnitFloor(from, ba.internalRes)
	toUnit := av.TimeToUnitFloor(to, ba.internalRes)
	ba.setUnitInternal(fromUnit, toUnit, value)
}

func (ba *SegmentBitAv) Get(from, to time.Time, res av.TimeResolution) *BitVector {

	if res > ba.internalRes {
		// lower resolution
		fromUnit := av.TimeToUnitFloor(av.FloorDate(from, res), ba.internalRes)
		toUnit := av.TimeToUnitFloor(av.CeilDate(to, res), ba.internalRes)
		arr := ba.getUnitInternal(fromUnit, toUnit)
		factor := int(res / ba.internalRes)
		reducedArr := reduceByFactor(arr, factor, reduceAllOne)
		return NewBitVector(res, ba.internalRes, reducedArr, av.FloorDate(from, res))

	} else if res < ba.internalRes {
		// higher resolution
		fromUnitInternalRes := av.TimeToUnitFloor(from, ba.internalRes)
		toUnitInternalRes := av.TimeToUnitFloor(av.CeilDate(to, ba.internalRes), ba.internalRes)
		arr := ba.getUnitInternal(fromUnitInternalRes, toUnitInternalRes)
		factor := int(ba.internalRes / res)
		arrMultiplied := multiplyByFactor(arr, factor)
		cutoff := av.TimeToUnitFloor(from, res) - fromUnitInternalRes*factor
		origlen := av.TimeToUnitFloor(to, res) - av.TimeToUnitFloor(from, res)
		arrTrimmed := arrMultiplied[cutoff : cutoff+origlen]
		return NewBitVector(res, ba.internalRes, arrTrimmed, av.FloorDate(from, ba.internalRes))
	} else {
		// internal resolution
		fromUnit := av.TimeToUnitFloor(from, res)
		toUnit := av.TimeToUnitFloor(to, res)
		arr := ba.getUnitInternal(fromUnit, toUnit)
		return NewBitVector(res, ba.internalRes, arr, av.FloorDate(from, res))
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

func (ba *SegmentBitAv) SetAt(at time.Time, value byte) {
	atUnit := av.TimeToUnitFloor(at, ba.internalRes)
	ba.setUnitInternal(atUnit, atUnit+1, value)
}

func (ba *SegmentBitAv) GetAt(at time.Time) byte {
	atUnit := av.TimeToUnitFloor(at, ba.internalRes)
	arr := ba.getUnitInternal(atUnit, atUnit+1)
	return byte(arr[0])
}

func (ba *SegmentBitAv) String() string {
	var buffer bytes.Buffer

	segments := ba.bsp.FindAll(ba.ID)
	for _, segment := range segments {
		buffer.WriteString(strconv.Itoa(segment.start))
		buffer.WriteString("->")
		buffer.WriteString(segment.String())
		buffer.WriteRune('\n')
	}

	return buffer.String()
}

func (ba *SegmentBitAv) segmentStart(i int) int {
	return i - i%ba.segmentSize
}

func (ba *SegmentBitAv) getOrEmptyBitSegment(startValue int) *BitSegment {
	if segment := ba.bsp.Find(ba.ID, startValue); segment != nil {
		return segment
	}
	return NewBitSegment(ba.ID, startValue)
}

func (ba *SegmentBitAv) setUnitInternal(from, to int, value byte) {
	currentBitSegment := ba.getOrEmptyBitSegment(ba.segmentStart(from))
	for i, j := from, from%ba.segmentSize; i < to; i, j = i+1, j+1 {
		if j == ba.segmentSize {
			ba.bsp.Save(currentBitSegment)
			currentBitSegment = ba.getOrEmptyBitSegment(i)
			j = 0
		}
		currentBitSegment.SetBit(&currentBitSegment.Int, j, uint(value))
	}
	ba.bsp.Save(currentBitSegment)
}

func (ba *SegmentBitAv) getUnitInternal(from, to int) []byte {
	length := to - from
	result := make([]byte, length)
	currentBitSegment := ba.getOrEmptyBitSegment(ba.segmentStart(from))
	for i, j := 0, from%ba.segmentSize; i < length; i, j = i+1, j+1 {
		if j == ba.segmentSize {
			currentBitSegment = ba.getOrEmptyBitSegment(i + from)
			j = 0
		}
		result[i] = byte(currentBitSegment.Bit(j))
	}
	return result
}
